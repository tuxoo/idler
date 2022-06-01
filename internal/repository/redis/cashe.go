package redis

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
)

type Cache interface {
	Set(ctx context.Context, key string, value *entity.User)
	Get(ctx context.Context, key string) *entity.User
}
