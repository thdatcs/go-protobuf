package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"go-protobuf/api"
	"go-protobuf/helpers"
	"go-protobuf/repositories/caches"
	"go-protobuf/repositories/databases"
	"go-protobuf/repositories/queues"
	"go-protobuf/servers"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	_ "go-protobuf/cmd/statik"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
)

func init() {
	if err := helpers.InitZap(); err != nil {
		return
	}
}

func main() {
	configFile := flag.String("configFile", "", "Path of the configuration file")
	flag.Parse()

	config, err := helpers.InitConfig(*configFile)
	if err != nil {
		zap.S().Errorw("Could not init config", zap.Error(err))
		return
	}
	statik, err := fs.New()
	if err != nil {
		zap.S().Errorw("Could not create statik", zap.Error(err))
		return
	}
	mysql, err := helpers.InitMySQL(config.MySQL.Host, config.MySQL.Port,
		config.MySQL.Username, config.MySQL.Password, config.MySQL.Database)
	if err != nil {
		zap.S().Errorw("Could not create instance mysql", zap.Error(err))
		return
	}
	zap.S().Info("MySQL connected")
	asyncProducers := make(map[string]sarama.AsyncProducer)
	topics := make(map[string][]string)
	for key, value := range config.KafkaProducers {
		asyncProducer, err := helpers.InitKafkaAsyncProducer(value.Version, value.Brokers)
		if err != nil {
			zap.S().Errorw("Could not create instance kafka async producer", zap.Error(err))
			//return
		}
		asyncProducers[key] = asyncProducer
		topics[key] = value.Topics
	}
	zap.S().Info("Kafka connected")
	redis, err := helpers.InitRedis(config.Redis.Host, config.Redis.Port,
		config.Redis.Password, config.Redis.DB)
	if err != nil {
		zap.S().Errorw("Could not create instance redis", zap.Error(err))
		//return
	}
	zap.S().Info("Redis connected")
	_, err = helpers.InitOpentracing(config.Jaeger.Host, config.Jaeger.Port, config.Jaeger.Name)
	if err != nil {
		zap.S().Errorw("Could not create instance jaeger", zap.Error(err))
		// return
	}
	zap.S().Info("Jaeger connected")
	serverAddr := fmt.Sprintf("%s:%d", "localhost", config.ServerPort)
	gatewayAddr := fmt.Sprintf("%s:%d", config.GatewayHost, config.GatewayPort)

	// Init adapter
	// Init cache
	redisCache := caches.NewRedisCache(redis)
	// Init queue
	kafkaQueue := queues.NewKafkaQueue(asyncProducers, topics)
	// Init database
	defaultDB := databases.NewDefaultDB(mysql)
	userDatabase := databases.NewUserDatabase()

	// Init server
	pingServer := servers.NewPingServer()
	userServer := servers.NewUserServer(config, defaultDB, redisCache, kafkaQueue, userDatabase)

	go func() {
		listener, err := net.Listen("tcp", serverAddr)
		if err != nil {
			zap.S().Errorw("Server: Failed to listen port %d", zap.Error(err))
			return
		}

		//opentracing.SetGlobalTracer(tracer)
		server := grpc.NewServer(grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(),
			)),
		)

		// Register the PingServer
		api.RegisterPingServer(server, pingServer)
		// Register the UserServer
		api.RegisterUserServer(server, userServer)

		if err := server.Serve(listener); err != nil {
			zap.S().Errorw("Server: Failed to serve", zap.Error(err))
			return
		}
	}()
	zap.S().Infof("Server is started on port %d", config.ServerPort)

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		runtime.HTTPError = func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
			w.Header().Set("Content-Type", marshaler.ContentType())
			w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))
			if err := json.NewEncoder(w).Encode(status.Convert(err).Message()); err != nil {
				zap.S().Errorw("Gateway: Failed to write response", zap.Error(err))
			}
		}

		gwmux := runtime.NewServeMux()

		opts := []grpc.DialOption{grpc.WithInsecure()}

		// Register the PingHandler
		err = api.RegisterPingHandlerFromEndpoint(ctx, gwmux, serverAddr, opts)
		if err != nil {
			zap.S().Errorw("Gateway: Failed to register PingHandler", zap.Error(err))
			return
		}
		// Register the UserHandler
		err = api.RegisterUserHandlerFromEndpoint(ctx, gwmux, serverAddr, opts)
		if err != nil {
			zap.S().Errorw("Gateway: Failed to register UserHandler", zap.Error(err))
			return
		}

		mux := http.NewServeMux()
		mux.Handle("/", gwmux)
		mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(statik)))

		gateway := &http.Server{
			Addr:    gatewayAddr,
			Handler: mux,
		}

		if err := gateway.ListenAndServe(); err != nil {
			zap.S().Errorw("Gateway: Failed to listen and serve", zap.Error(err))
			return
		}
	}()
	zap.S().Infof("Gateway is started on port %d", config.GatewayPort)

	zap.S().Info("*****RUNNING******")

	signals := make(chan os.Signal, 1)
	shutdown := make(chan bool, 1)

	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals

		// TODO: Release resources
		// thalesClient.Close()

		shutdown <- true
	}()
	<-shutdown

	zap.S().Info("*****SHUTDOWN*****")
}
