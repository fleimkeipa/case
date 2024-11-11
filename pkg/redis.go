package pkg

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return redisClient
}
