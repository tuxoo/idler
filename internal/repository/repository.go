package repository

import (
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	SaveUser(user entity.User) (int, error)
	GetUser(name, email, password string) (entity.User, error)
}

type Repositories struct {
	Users Users
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users: postgres.NewUserRepository(db),
	}
}
