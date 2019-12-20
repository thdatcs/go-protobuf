package configs

// Config represents all configurations
type Config struct {
	ServerPort          int
	GatewayHost         string
	GatewayPort         int
	MySQL               mysql
	Redis               redis
	Jaeger              jaeger
	KafkaConsumer       kafkaConsumer
	KafkaProducers      map[string]kafkaProducer
	RabbitMQ            rabbitmq
}

type mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type jaeger struct {
	Host string
	Port int
	Name string
}

type kafkaConsumer struct {
	Version string
	Brokers []string
	Topics  []string
	GroupID string
}

type kafkaProducer struct {
	Version string
	Brokers []string
	Topics  []string
}

type rabbitmq struct {
	Host string
	Port int
	Name string
}
