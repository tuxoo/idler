package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/internal/repository/redis"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"time"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (auth.Token, error)
	GetAll() ([]entity.User, error)
	//GenerateToken(username, password string) (string, error)
	//ParseToken(token string) (int, error)
}

type Services struct {
	Users Users
}

type ServicesDepends struct {
	Repositories *repository.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	TokenTTL     time.Duration
	UserCache    redis.Cache
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.Repositories.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL, deps.UserCache)

	return &Services{
		Users: userService,
	}
}
