package postgres_repository

import (
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	. "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user entity.User) (*dto.UserDTO, error) {
	var newUser dto.UserDTO
	query := fmt.Sprintf("INSERT INTO %s (name, login_email, password_hash, registered_at, visited_at, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, login_email, registered_at, visited_at, role, is_enabled", userTable)
	row := r.db.QueryRowx(query, user.Name, user.LoginEmail, user.PasswordHash, user.RegisteredAt, user.VisitedAt, entity.UserRole)

	if err := row.StructScan(&newUser); err != nil {
		return &newUser, err
	}

	return &newUser, nil
}

func (r *UserRepository) UpdateById(id UUID) error {
	query := fmt.Sprintf("UPDATE %s SET is_enabled=true WHERE id=$1", userTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) FindByCredentials(email, password string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true AND login_email=$1 AND password_hash=$2", userTable)
	if err := r.db.Get(&user, query, email, password); err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *UserRepository) FindById(id UUID) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true AND id=$1", userTable)
	if err := r.db.Get(&user, query, id); err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *UserRepository) FindAll() ([]dto.UserDTO, error) {
	var users []dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true", userTable)
	if err := r.db.Select(&users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (*dto.UserDTO, error) {
	var user dto.UserDTO
	query := fmt.Sprintf("SELECT id, name, login_email, registered_at, visited_at, role, is_enabled FROM %s WHERE is_enabled=true AND WHERE login_email=$1", userTable)
	if err := r.db.Get(&user, query, email); err != nil {
		return &user, err
	}

	return &user, nil
}
