package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"time"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (auth.Token, error)
	GetAll() ([]entity.User, error)
}

type Services struct {
	UserService Users
}

type ServicesDepends struct {
	Repositories *repository.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	TokenTTL     time.Duration
	UserCache    repository.UserCache
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.Repositories.UserRepository, deps.Hasher, deps.TokenManager, deps.TokenTTL, deps.UserCache)

	return &Services{
		UserService: userService,
	}
}
