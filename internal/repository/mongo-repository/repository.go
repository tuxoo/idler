package mongo_repository

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	. "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	messageCollection = "message"
)

type Messages interface {
	Save(ctx context.Context, message entity.Message) error
	SaveAll(ctx context.Context, messages []entity.Message) error
	FindByConversationId(ctx context.Context, conversationId UUID) (entity.Message, error)
	FindAllByConversationId(ctx context.Context, conversationId UUID) ([]entity.Message, error)
}

type Repositories struct {
	Messages Messages
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Messages: NewMessageRepository(db),
	}
}
