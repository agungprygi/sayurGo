package config

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(cfg *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}
