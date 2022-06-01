package service

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"time"
)

type Users interface {
	SignUp(user dto.SignUpDTO) error
	SignIn(user dto.SignInDTO) (auth.Token, error)
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
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.Repositories.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL)

	return &Services{
		Users: userService,
	}
}
