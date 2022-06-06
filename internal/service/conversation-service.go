package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres-repositrory"
	"github.com/gin-gonic/gin"
)

type ConversationService struct {
	repository postgres_repositrory.Conversations
}

func NewConversationService(repository postgres_repositrory.Conversations) *ConversationService {
	return &ConversationService{
		repository: repository,
	}
}

func (s *ConversationService) CreateConversation(ctx context.Context, userId int, conversationDTO dto.ConversationDTO) error {
	conversation := entity.Conversation{
		Name:  conversationDTO.Name,
		Owner: userId,
		//Participants: []dto.UserDTO{conversationDTO.Participant},
	}

	_, err := s.repository.Save(conversation)

	return err
}

func (s *ConversationService) GetAll(c *gin.Context) ([]dto.ConversationDTO, error) {
	return s.repository.FindAll()
}

func (s *ConversationService) GetById(ctx context.Context, id int) (*dto.ConversationDTO, error) {
	return s.repository.FindById(id)
}

func (s *ConversationService) RemoveById(ctx context.Context, id int) error {
	return s.repository.DeleteById(id)
}
