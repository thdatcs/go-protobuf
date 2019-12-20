package queues

import (
	"context"

	"github.com/golang/protobuf/proto"
)

// Queue is repository of queue
type Queue interface {
	SendProtoAsync(ctx context.Context, producer, topic, key string, value proto.Message) error
}
