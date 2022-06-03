package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"strconv"
	"time"
)

type UserService struct {
	repository   repository.UserRepository
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	tokenTTL     time.Duration
	userCache    cache.Cache[string, entity.User]
}

func NewUserService(repository repository.UserRepository, hasher hash.PasswordHasher, tokenManager auth.TokenManager, tokenTTL time.Duration, userCache cache.Cache[string, entity.User]) *UserService {
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

	id, err := s.repository.Save(user)
	user.Id = id

	s.userCache.Set(ctx, dto.Email, &user)
	return err
}

func (s *UserService) SignIn(ctx context.Context, dto dto.SignInDTO) (auth.Token, error) {
	var userPointer, err = s.userCache.Get(ctx, dto.Email)
	if err != nil {
		return "", err
	}

	if userPointer == nil {
		user, err := s.repository.FindByCredentials(dto.Email, s.hasher.Hash(dto.Password))
		if err != nil {
			return "", err
		}

		s.userCache.Set(ctx, dto.Email, &user)
		return s.tokenManager.GenerateToken(strconv.Itoa(user.Id), s.tokenTTL)
	}
	return s.tokenManager.GenerateToken(strconv.Itoa(userPointer.Id), s.tokenTTL)
}

func (s *UserService) GetAll(ctx context.Context) ([]entity.User, error) {
	return s.repository.FindAll()
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.repository.FindByEmail(email)
}
