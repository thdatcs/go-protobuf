package caches

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
)

// Cache is repository of cache
type Cache interface {
	Exists(ctx context.Context, key string) error
	GetInt(ctx context.Context, key string, del bool) (int, error)
	GetString(ctx context.Context, key string, del bool) (string, error)
	GetBytes(ctx context.Context, key string, del bool) ([]byte, error)
	GetProto(ctx context.Context, key string, value proto.Message, del bool) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetProto(ctx context.Context, key string, value proto.Message, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	SetTx(ctx context.Context, data map[string]interface{}, expiration time.Duration) error
	SetProtoTx(ctx context.Context, data map[string]proto.Message, expiration time.Duration) error
	DelTx(ctx context.Context, keys ...string) error
}
