package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"time"
)

type Users interface {
	RegisterUser(ctx context.Context, user dto.SignUpDTO) error
	AuthorizeUser(ctx context.Context, user dto.SignUpDTO) (auth.Token, error)
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
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.Repositories.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL)

	return &Services{
		Users: userService,
	}
}
