package redisinfra

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
	Ctx    = context.Background()
)

func GetRedisClient() *redis.Client {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr: "192.168.1.233:6379",
			DB:   0,	
		})
	})
	return client
}
