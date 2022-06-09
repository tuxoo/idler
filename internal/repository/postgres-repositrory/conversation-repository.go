package postgres_repositrory

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	. "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type ConversationRepository struct {
	db *sqlx.DB
}

func NewConversationRepository(db *sqlx.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) Save(conversation entity.Conversation) (*dto.ConversationDTO, error) {
	var newConversation dto.ConversationDTO

	query := fmt.Sprintf("INSERT INTO %s (name, owner) VALUES ($1, $2) RETURNING name, owner", conversationTable)
	row := r.db.QueryRowx(query, conversation.Name, conversation.Owner)
	if err := row.StructScan(&newConversation); err != nil {
		time.Now().Format(time.Layout)
		return &newConversation, err
	}

	return &newConversation, nil
}

func (r *ConversationRepository) FindAll() ([]dto.ConversationDTO, error) {
	var conversations []dto.ConversationDTO
	query := fmt.Sprintf("SELECT name, owner FROM %s", conversationTable)

	err := r.db.Select(&conversations, query)
	if err != nil {
		return conversations, err
	}

	return conversations, nil
}

func (r *ConversationRepository) FindById(id UUID) (*dto.ConversationDTO, error) {
	var conversation dto.ConversationDTO
	query := fmt.Sprintf("SELECT name, owner FROM %s WHERE id=$1", conversationTable)

	row := r.db.QueryRowx(query, id)
	if err := row.StructScan(&conversation); err != nil { // TODO: Perhaps need convert id to string
		return &conversation, err
	}

	return &conversation, nil
}

func (r *ConversationRepository) DeleteById(id UUID) error {
	query := fmt.Sprintf("DELETE FROM %s where id=$1", conversationTable)
	_, err := r.db.Exec(query, id)
	return err
}
