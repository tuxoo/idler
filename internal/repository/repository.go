package repository

import (
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	SaveUser(user entity.User) (int, error)
	GetUser(email, password string) (entity.User, error)
	GetAll() ([]entity.User, error)
}

type Repositories struct {
	UserRepository UserRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository: postgres.NewUserRepository(db),
	}
}
