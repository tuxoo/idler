package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres-repositrory"
	. "github.com/google/uuid"
)

type ConversationService struct {
	repository postgres_repository.Conversations
}

func NewConversationService(repository postgres_repository.Conversations) *ConversationService {
	return &ConversationService{
		repository: repository,
	}
}

func (s *ConversationService) CreateConversation(ctx context.Context, userId UUID, conversationDTO dto.ConversationDTO) error {
	conversation := entity.Conversation{
		Name:  conversationDTO.Name,
		Owner: userId,
		//Participants: []dto.UserDTO{conversationDTO.Participant},
	}

	_, err := s.repository.Save(ctx, conversation)

	return err
}

func (s *ConversationService) GetByOwnerId(ctx context.Context, id UUID) ([]dto.ConversationDTO, error) {
	return s.repository.FindByOwnerId(ctx, id)
}

func (s *ConversationService) GetById(ctx context.Context, id UUID) (*dto.ConversationDTO, error) {
	return s.repository.FindById(ctx, id)
}

func (s *ConversationService) RemoveById(ctx context.Context, id UUID) error {
	return s.repository.DeleteById(ctx, id)
}
