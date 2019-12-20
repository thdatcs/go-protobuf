package queues

import (
	"go-protobuf/utils/jaeger"
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
)

type kafkaQueue struct {
	asyncProducers map[string]sarama.AsyncProducer
	topics         map[string][]string
}

// NewKafkaQueue creates a new instance
func NewKafkaQueue(asyncProducers map[string]sarama.AsyncProducer, topics map[string][]string) Queue {
	return &kafkaQueue{
		asyncProducers: asyncProducers,
		topics:         topics,
	}
}

func (q *kafkaQueue) SendProtoAsync(ctx context.Context, producer, topic, key string, value proto.Message) (err error) {
	span := jaeger.Start(ctx, ">queues.kafkaQueue/SendProtoAsync", ext.SpanKindProducer)
	defer func() {
		jaeger.Finish(span, err)
	}()

	var (
		asyncProducer sarama.AsyncProducer
		buffer        []byte
	)
	asyncProducer, err = q.validateProducer(producer)
	if err != nil {
		return err
	}
	err = q.validateTopic(producer, topic)
	if err != nil {
		return err
	}

	buffer, err = proto.Marshal(value)
	if err != nil {
		return err
	}
	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(buffer),
	}

	if err := jaeger.InjectKafkaHeaders(span, &message.Headers); err != nil {
		zap.S().Warnw(fmt.Sprintf("%v-Failed to inject span", message.Key), zap.Error(err))
	}

	asyncProducer.Input() <- message
	return nil
}

func (q *kafkaQueue) validateProducer(producer string) (sarama.AsyncProducer, error) {
	asyncProducer, ok := q.asyncProducers[producer]
	if !ok {
		return nil, fmt.Errorf("Invaild producer %v", producer)
	}
	return asyncProducer, nil
}

func (q *kafkaQueue) validateTopic(producer, topic string) error {
	var flag bool
	array := q.topics[producer]
	for _, item := range array {
		if item == topic {
			flag = !flag
			break
		}
	}
	if !flag {
		return fmt.Errorf("Invalid topic %v in producer %v", topic, producer)
	}
	return nil
}
