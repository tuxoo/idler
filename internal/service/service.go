package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"github.com/gin-gonic/gin"
	"time"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (auth.Token, error)
	GetById(ctx context.Context, id int) (*dto.UserDTO, error)
	GetAll(ctx context.Context) ([]dto.UserDTO, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
}

type Conversations interface {
	CreateConversation(ctx context.Context, userId int, conversation dto.ConversationDTO) error
	GetAll(c *gin.Context) ([]dto.ConversationDTO, error)
	GetById(ctx context.Context, id int) (*dto.ConversationDTO, error)
	RemoveById(ctx context.Context, id int) error
}

type Services struct {
	UserService         Users
	ConversationService Conversations
}

type ServicesDepends struct {
	Repositories *postgres.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	TokenTTL     time.Duration
	UserCache    cache.Cache[string, dto.UserDTO]
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.Repositories.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL, deps.UserCache)
	conversationService := NewConversationService(deps.Repositories.Conversations)

	return &Services{
		UserService:         userService,
		ConversationService: conversationService,
	}
}
