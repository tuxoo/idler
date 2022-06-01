package redis

import (
	"context"
	json2 "encoding/json"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/go-redis/redis/v8"
	"time"
)

type UserCache struct {
	RedisClient *redis.Client
	Expires     time.Duration
}

func NewUserCache(client *redis.Client, expires time.Duration) repository.UserCache {
	return &UserCache{
		RedisClient: client,
		Expires:     expires,
	}
}

func (c *UserCache) Set(ctx context.Context, key string, value *entity.User) {
	json, err := json2.Marshal(value)
	if err != nil {
		panic(err)
	}

	c.RedisClient.Set(ctx, key, json, c.Expires*time.Second)
}

func (c *UserCache) Get(ctx context.Context, key string) *entity.User {
	value, err := c.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	user := entity.User{}
	if err := json2.Unmarshal([]byte(value), &user); err != nil {
		panic(err)
	}

	return &user
}
