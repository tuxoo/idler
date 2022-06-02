package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/gin-gonic/gin"
	"time"
)

type DialogService struct {
	repository repository.DialogRepository
}

func NewDialogService(repository repository.DialogRepository) *DialogService {
	return &DialogService{
		repository: repository,
	}
}

func (s *DialogService) CreateDialog(ctx context.Context, dto dto.DialogDTO) error {
	dialog := entity.Dialog{
		Name:         dto.Name,
		CreatedAt:    time.Now(),
		LastMessage:  time.Now(),
		FirstUserId:  dto.FirstId,
		SecondUserId: dto.SecondId,
	}

	_, err := s.repository.Save(dialog)

	return err
}

func (s *DialogService) GetAll(c *gin.Context) ([]entity.Dialog, error) {
	return s.repository.FindAll()
}

func (s *DialogService) GetById(ctx context.Context, id int) (entity.Dialog, error) {
	return s.repository.FindById(id)
}

func (s *DialogService) RemoveById(ctx context.Context, id int) error {
	return s.repository.DeleteById(id)
}
