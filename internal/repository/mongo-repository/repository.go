package mongo_repository

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	messageCollection = "message"
)

type Messages interface {
	Save(ctx context.Context, message entity.Message) error
}

type Repositories struct {
	Messages Messages
}

func NewRepositories(db *mongo.Database) *Repositories {
	return &Repositories{
		Messages: NewMessageRepositories(db),
	}
}
