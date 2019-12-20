package helpers

import (
	"fmt"

	"github.com/go-redis/redis"
)

// InitRedis initializes redis standalone
func InitRedis(host string, port int, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", host, port),
		Password: password,
		DB:       db,
	})
	_, err := client.Ping().Result()
	return client, err
}
