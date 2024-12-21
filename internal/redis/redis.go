package redis

import (
	"context"
	"fmt"
	"project/config"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

// InitializeRedis sets up the Redis connection
func InitializeRedis(cfg *config.Config) error {
	Redis = redis.NewClient(&redis.Options{
		Addr: cfg.Redis.URI,
	})
	if _, err := Redis.Ping(context.Background()).Result(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return nil
}

// CloseRedis closes the Redis connection
func CloseRedis() {
	if Redis != nil {
		_ = Redis.Close()
	}
}
