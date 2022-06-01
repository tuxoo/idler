package redis

import (
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	Address  string
	DB       int
	Password string
	Expires  time.Duration
}

func NewRedisCache(cfg config.RedisConfig) Cache {
	return &RedisCache{
		Address:  fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		DB:       cfg.DB,
		Password: cfg.Password,
		Expires:  cfg.Expires,
	}
}

func (c *RedisCache) getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,
	})
}

func (c *RedisCache) Set(ctx context.Context, key string, value *entity.User) {
	client := c.getRedisClient()

	json, err := json2.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(ctx, key, json, c.Expires*time.Second)
}

func (c *RedisCache) Get(ctx context.Context, key string) *entity.User {
	client := c.getRedisClient()

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	user := entity.User{}
	if err := json2.Unmarshal([]byte(value), &user); err != nil {
		panic(err)
	}

	return &user
}
