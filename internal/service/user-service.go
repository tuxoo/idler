package service

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"time"
)

type UserService struct {
	repository   repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	tokenTTL     time.Duration
}

func NewUserService(repository repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, tokenTTL time.Duration) *UserService {
	return &UserService{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
		tokenTTL:     tokenTTL,
	}
}

func (s *UserService) SignUp(dto dto.SignUpDTO) error {
	user := entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     s.hasher.Hash(dto.Password),
		RegisteredAt: time.Now(),
		VisitedAt:    time.Now(),
	}
	_, err := s.repository.SaveUser(user)
	return err
}

func (s *UserService) SignIn(dto dto.SignInDTO) (auth.Token, error) {
	user, err := s.repository.GetUser(dto.Email, s.hasher.Hash(dto.Password))
	if err != nil {
		return "", err
	}

	return s.tokenManager.GenerateToken(string(user.Id), s.tokenTTL)
}

func (s *UserService) GetAll() ([]entity.User, error) {
	return s.repository.GetAll()
}
