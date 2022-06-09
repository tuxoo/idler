package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	mongo_repository "github.com/eugene-krivtsov/idler/internal/repository/mongo-repository"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres-repositrory"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"github.com/gin-gonic/gin"
	. "github.com/google/uuid"
	"time"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	VerifyUser(ctx context.Context, code UUID) error
	SignIn(ctx context.Context, user dto.SignInDTO) (auth.Token, error)
	GetById(ctx context.Context, id UUID) (*dto.UserDTO, error)
	GetAll(ctx context.Context) ([]dto.UserDTO, error)
	GetByEmail(ctx context.Context, email string) (*dto.UserDTO, error)
}

type Conversations interface {
	CreateConversation(ctx context.Context, userId UUID, conversation dto.ConversationDTO) error
	GetAll(c *gin.Context) ([]dto.ConversationDTO, error)
	GetById(ctx context.Context, id UUID) (*dto.ConversationDTO, error)
	RemoveById(ctx context.Context, id UUID) error
}

type Messages interface {
	Create(ctx context.Context, message entity.Message) error
	CreateAll(ctx context.Context, messages []entity.Message) error
	GetByConversationId(ctx context.Context, id UUID) (entity.Message, error)
	GetAllConversationId(ctx context.Context, id UUID) ([]entity.Message, error)
}

type Services struct {
	UserService         Users
	ConversationService Conversations
	MessageService      Messages
}

type ServicesDepends struct {
	PostgresRepositories *postgres_repository.Repositories
	MongoRepositories    *mongo_repository.Repositories
	Hasher               hash.PasswordHasher
	TokenManager         auth.TokenManager
	TokenTTL             time.Duration
	UserCache            cache.Cache[string, dto.UserDTO]
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.PostgresRepositories.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL, deps.UserCache)
	conversationService := NewConversationService(deps.PostgresRepositories.Conversations)
	messageService := NewMessageService(deps.MongoRepositories.Messages)

	return &Services{
		UserService:         userService,
		ConversationService: conversationService,
		MessageService:      messageService,
	}
}
