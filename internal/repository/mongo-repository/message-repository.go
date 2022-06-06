package mongo_repository

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	db *mongo.Collection
}

func NewMessageRepositories(db *mongo.Database) *MessageRepository {
	return &MessageRepository{
		db: db.Collection(messageCollection),
	}
}

func (r *MessageRepository) Save(ctx context.Context, message entity.Message) error {
	_, err := r.db.InsertOne(ctx, message)
	return err
}
