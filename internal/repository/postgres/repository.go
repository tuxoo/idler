package postgres

import (
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	Save(user entity.User) (int, error)
	FindByCredentials(email, password string) (entity.User, error)
	FindAll() ([]entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type Dialogs interface {
	Save(dialog entity.Dialog) (int, error)
	FindAll() ([]entity.Dialog, error)
	FindById(id int) (entity.Dialog, error)
	DeleteById(id int) error
}

type Repositories struct {
	Users   Users
	Dialogs Dialogs
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users:   NewUserRepository(db),
		Dialogs: NewDialogRepository(db),
	}
}
