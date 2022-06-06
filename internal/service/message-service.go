package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/mongo-repository"
)

type MessageService struct {
	repository mongo_repository.Messages
}

func NewMessageService(repository mongo_repository.Messages) *MessageService {
	return &MessageService{
		repository: repository,
	}
}

func (s *MessageService) Save(ctx context.Context, message entity.Message) error {
	return s.repository.Save(ctx, message)
}
