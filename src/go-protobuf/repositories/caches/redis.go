package caches

import (
	"go-protobuf/utils/jaeger"
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/go-redis/redis"
)

type redisCache struct {
	client *redis.Client
}

// NewRedisCache creates a new instance
func NewRedisCache(client *redis.Client) Cache {
	return &redisCache{
		client: client,
	}
}

func (c *redisCache) Exists(ctx context.Context, key string) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/Exists", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	indicator, err := c.client.Exists(key).Result()
	if err != nil {
		return err
	}
	if indicator == 0 {
		return redis.Nil
	}
	return nil
}

func (c *redisCache) GetInt(ctx context.Context, key string, del bool) (value int, err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/GetInt", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	value, err = c.client.Get(key).Int()
	if err != nil {
		return 0, err
	}

	if del {
		_ = c.Del(ctx, key)
	}

	return value, nil
}

func (c *redisCache) GetString(ctx context.Context, key string, del bool) (value string, err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/GetString", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	value, err = c.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	if del {
		_ = c.Del(ctx, key)
	}

	return value, nil
}

func (c *redisCache) GetBytes(ctx context.Context, key string, del bool) (value []byte, err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/GetBytes", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	value, err = c.client.Get(key).Bytes()
	if err != nil {
		return nil, err
	}

	if del {
		_ = c.Del(ctx, key)
	}

	return value, nil
}

func (c *redisCache) GetProto(ctx context.Context, key string, value proto.Message, del bool) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/GetProto", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	data, err := c.GetBytes(ctx, key, del)
	if err != nil {
		return err
	}
	err = proto.Unmarshal(data, value)
	if err != nil {
		return err
	}

	return nil
}

func (c *redisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/Set", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	_, err = c.client.Set(key, value, expiration).Result()
	return err
}

func (c *redisCache) SetProto(ctx context.Context, key string, value proto.Message, expiration time.Duration) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/SetProto", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	data, err := proto.Marshal(value)
	if err != nil {
		return err
	}
	_, err = c.client.Set(key, data, expiration).Result()
	return err
}

func (c *redisCache) Del(ctx context.Context, key string) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/Del", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	_, err = c.client.Del(key).Result()
	return err
}

func (c *redisCache) Incr(ctx context.Context, key string) (value int64, err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/Incr", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	return c.client.Incr(key).Result()
}

func (c *redisCache) SetTx(ctx context.Context, data map[string]interface{}, expiration time.Duration) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/SetTx", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	txFunc := func(tx *redis.Tx) error {
		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			for k, v := range data {
				_, err := c.client.Set(k, v, expiration).Result()
				if err != nil {
					return err
				}
			}
			return nil
		})
		return err
	}
	return c.client.Watch(txFunc)
}

func (c *redisCache) SetProtoTx(ctx context.Context, data map[string]proto.Message, expiration time.Duration) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/SetProtoTx", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	txFunc := func(tx *redis.Tx) error {
		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			for k, v := range data {
				err := c.SetProto(ctx, k, v, expiration)
				if err != nil {
					return err
				}
			}
			return nil
		})
		return err
	}
	return c.client.Watch(txFunc)
}

func (c *redisCache) DelTx(ctx context.Context, keys ...string) (err error) {
	span := jaeger.Start(ctx, ">caches.redisCache/DelTx", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	txFunc := func(tx *redis.Tx) error {
		_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
			for _, k := range keys {
				err := c.Del(ctx, k)
				if err != nil {
					return err
				}
			}
			return nil
		})
		return err
	}
	return c.client.Watch(txFunc)
}
