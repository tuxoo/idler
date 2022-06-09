package postgres_repositrory

import (
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	. "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	userTable         = "\"user\""
	conversationTable = "conversation"
)

type Users interface {
	Save(user entity.User) (*dto.UserDTO, error)
	FindByCredentials(email, password string) (*dto.UserDTO, error)
	FindById(id UUID) (*dto.UserDTO, error)
	FindAll() ([]dto.UserDTO, error)
	FindByEmail(email string) (*dto.UserDTO, error)
}

type Conversations interface {
	Save(conversation entity.Conversation) (*dto.ConversationDTO, error)
	FindAll() ([]dto.ConversationDTO, error)
	FindById(id UUID) (*dto.ConversationDTO, error)
	DeleteById(id UUID) error
}

type Repositories struct {
	Users         Users
	Conversations Conversations
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users:         NewUserRepository(db),
		Conversations: NewConversationRepository(db),
	}
}
