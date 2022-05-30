package postgres

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model"
	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s.%s (name, username, password_hash) values ($1, $2, $3) RETURNING id", schema, usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UsersRepository) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s.%s WHERE username=$1 AND password_hash=$2", schema, usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
