package redis

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
