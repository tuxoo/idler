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

type Messages interface {
	Save(ctx context.Context, message entity.Message) error
}

type Services struct {
	UserService         Users
	ConversationService Conversations
	MessageService      Messages
}

type ServicesDepends struct {
	PostgresRepositories *postgres_repositrory.Repositories
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
