package service

import (
	"github.com/eugene-krivtsov/idler/internal/model"
	"github.com/eugene-krivtsov/idler/internal/repository"
)

type Users interface {
	CreateUser(user model.User) (int, error)
	//GenerateToken(username, password string) (string, error)
	//ParseToken(token string) (int, error)
}

type Services struct {
	Users Users
}

type ServicesDepends struct {
	Repositories *repository.Repositories
}

func NewServices(deps ServicesDepends) *Services {
	return &Services{
		Users: NewUserService(deps.Repositories.Users),
	}
}
