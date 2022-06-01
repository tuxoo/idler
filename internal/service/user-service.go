package service

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/internal/repository/redis"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"strconv"
	"time"
)

type UserService struct {
	repository   repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	tokenTTL     time.Duration
	userCache    redis.Cache
}

func NewUserService(repository repository.Users, hasher hash.PasswordHasher, tokenManager auth.TokenManager, tokenTTL time.Duration, userCache redis.Cache) *UserService {
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

	_, err := s.repository.SaveUser(user)

	s.userCache.Set(ctx, dto.Email, &user)
	return err
}

func (s *UserService) SignIn(ctx context.Context, dto dto.SignInDTO) (auth.Token, error) {
	var userPointer = s.userCache.Get(ctx, dto.Email)
	if userPointer == nil {
		user, err := s.repository.GetUser(dto.Email, s.hasher.Hash(dto.Password))
		if err != nil {
			return "", err
		}

		fmt.Println("from db")
		s.userCache.Set(ctx, dto.Email, &user)
		return s.tokenManager.GenerateToken(strconv.Itoa(user.Id), s.tokenTTL)
	}
	fmt.Println("from cache")
	return s.tokenManager.GenerateToken(strconv.Itoa(userPointer.Id), s.tokenTTL)
}

func (s *UserService) GetAll() ([]entity.User, error) {
	return s.repository.GetAll()
}
