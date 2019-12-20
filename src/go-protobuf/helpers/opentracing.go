package helpers

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

// InitOpentracing initializes opentracing
func InitOpentracing(host string, port int, name string) (opentracing.Tracer, error) {
	cfg := config.Configuration{
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: fmt.Sprintf("%v:%v", host, port),
			LogSpans:           true,
		},
	}
	_, err := cfg.InitGlobalTracer(
		name,
		config.Logger(jaeger.StdLogger),
		config.Sampler(jaeger.NewConstSampler(true)),
		config.Metrics(metrics.NullFactory),
	)
	if err != nil {
		return nil, err
	}
	return opentracing.GlobalTracer(), nil
}
