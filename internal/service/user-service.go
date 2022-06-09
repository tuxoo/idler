package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres-repositrory"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	. "github.com/google/uuid"
	"time"
)

type UserService struct {
	repository   postgres_repositrory.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	tokenTTL     time.Duration
	userCache    cache.Cache[string, dto.UserDTO]
}

func NewUserService(repository postgres_repositrory.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, tokenTTL time.Duration, userCache cache.Cache[string, dto.UserDTO]) *UserService {
	return &UserService{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
		tokenTTL:     tokenTTL,
		userCache:    userCache,
	}
}

func (s *UserService) SignUp(ctx context.Context, dto dto.SignUpDTO) error {
	user := entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     s.hasher.Hash(dto.Password),
		RegisteredAt: time.Now(),
		VisitedAt:    time.Now(),
	}

	newUser, err := s.repository.Save(user)
	s.userCache.Set(ctx, dto.Email, newUser)
	return err
}

func (s *UserService) SignIn(ctx context.Context, dto dto.SignInDTO) (token auth.Token, err error) {
	user, err := s.userCache.Get(ctx, dto.Email)
	if err != nil {
		return
	}

	if user == nil {
		user, err := s.repository.FindByCredentials(dto.Email, s.hasher.Hash(dto.Password))
		if err != nil {
			return "", err
		}

		s.userCache.Set(ctx, dto.Email, user)

		token, err = s.tokenManager.GenerateToken(user.Id.String(), s.tokenTTL)
	}
	token, err = s.tokenManager.GenerateToken(user.Id.String(), s.tokenTTL)
	return
}

func (s *UserService) GetById(ctx context.Context, id UUID) (*dto.UserDTO, error) {
	return s.repository.FindById(id)
}

func (s *UserService) GetAll(ctx context.Context) ([]dto.UserDTO, error) {
	return s.repository.FindAll()
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error) {
	return s.repository.FindByEmail(email)
}
