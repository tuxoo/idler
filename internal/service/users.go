package service

import (
	"github.com/eugene-krivtsov/idler/internal/model"
	"github.com/eugene-krivtsov/idler/internal/repository"
)

type UserService struct {
	repository repository.Users
}

func NewUserService(repository repository.Users) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(user model.User) (int, error) {
	//user.Password = generatePasswordHash(user.Password) TODO:
	return s.repository.CreateUser(user)
}
