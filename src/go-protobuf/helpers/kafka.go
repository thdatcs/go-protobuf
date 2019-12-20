package helpers

import (
	"fmt"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

// InitKafkaSyncProducer initializes kafka sync producer
func InitKafkaSyncProducer(version string, brokers []string) (sarama.SyncProducer, error) {
	kafkaVersion, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return nil, err
	}
	config := sarama.NewConfig()
	config.Version = kafkaVersion
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	return producer, err
}

// InitKafkaAsyncProducer initializes kafka async producer
func InitKafkaAsyncProducer(version string, brokers []string) (sarama.AsyncProducer, error) {
	kafkaVersion, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return nil, err
	}
	config := sarama.NewConfig()
	config.Version = kafkaVersion
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err == nil {
		go func() {
			for {
				select {
				case producerErr := <-producer.Errors():
					zap.S().Errorw(fmt.Sprintf("%v-Failed to send to topic %v", producerErr.Msg.Key, producerErr.Msg.Topic), zap.Error(producerErr.Err))
					producer.Input() <- producerErr.Msg
				case message := <-producer.Successes():
					zap.S().Infof("%v-Sent to topic %v", message.Key, message.Topic)
				}
			}
		}()
	}
	return producer, err
}
