package repository

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
)

type UserCache interface {
	Set(ctx context.Context, key string, value *entity.User)
	Get(ctx context.Context, key string) *entity.User
}
