package postgres

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, registered_at, visited_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.RegisteredAt, user.VisitedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) FindByCredentials(email, password string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	if err := r.db.Get(&user, query, email, password); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	query := fmt.Sprintf("SELECT name, email FROM %s", usersTable)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1", usersTable)
	if err := r.db.Get(&user, query, email); err != nil {
		return user, err
	}

	return user, nil
}
