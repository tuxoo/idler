package postgres

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/jmoiron/sqlx"
)

type DialogRepository struct {
	db *sqlx.DB
}

func NewDialogRepository(db *sqlx.DB) *DialogRepository {
	return &DialogRepository{db: db}
}

func (r *DialogRepository) Save(dialog entity.Dialog) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, created_at, first_user_id, second_user_id) VALUES ($1, $2, $3, $4) RETURNING id", dialogsTable)
	row := r.db.QueryRow(query, dialog.Name, dialog.CreatedAt, dialog.FirstUserId, dialog.SecondUserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DialogRepository) FindAll() ([]entity.Dialog, error) {
	var dialogs []entity.Dialog
	query := fmt.Sprintf("SELECT name FROM %s", dialogsTable)
	if err := r.db.Get(&dialogs, query); err != nil {
		return nil, err
	}

	return dialogs, nil
}

func (r *DialogRepository) FindById(id int) (entity.Dialog, error) {
	var dialog entity.Dialog
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", dialogsTable)
	if err := r.db.Get(&dialog, query, id); err != nil { // TODO: Perhaps need convert id to string
		return dialog, err
	}

	return dialog, nil
}

func (r *DialogRepository) DeleteById(id int) error {
	query := fmt.Sprintf("U FROM %s where id=$1", dialogsTable)
	_, err := r.db.Exec(query, id)
	return err
}
