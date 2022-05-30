package repository

import (
	"github.com/eugene-krivtsov/idler/internal/model"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: postgres.NewUsersRepository(db),
	}
}
