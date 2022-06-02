package repository

import (
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Save(user entity.User) (int, error)
	FindByCredentials(email, password string) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type DialogRepository interface {
	Save(dialog entity.Dialog) (int, error)
	FindAll() ([]entity.Dialog, error)
	FindById(id int) (entity.Dialog, error)
	DeleteById(id int) error
}

type Repositories struct {
	UserRepository   UserRepository
	DialogRepository DialogRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository:   postgres.NewUserRepository(db),
		DialogRepository: postgres.NewDialogRepository(db),
	}
}
