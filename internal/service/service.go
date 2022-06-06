package service

import (
	"context"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	"github.com/gin-gonic/gin"
	"time"
)

type Users interface {
	SignUp(ctx context.Context, user dto.SignUpDTO) error
	SignIn(ctx context.Context, user dto.SignInDTO) (auth.Token, error)
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

type Dialogs interface {
	CreateDialog(ctx context.Context, user dto.DialogDTO) error
	GetAll(c *gin.Context) ([]entity.Dialog, error)
	GetById(ctx context.Context, id int) (entity.Dialog, error)
	RemoveById(ctx context.Context, id int) error
}

type Services struct {
	UserService   Users
	DialogService Dialogs
}

type ServicesDepends struct {
	Repositories *postgres.Repositories
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	TokenTTL     time.Duration
	UserCache    cache.Cache[string, entity.User]
}

func NewServices(deps ServicesDepends) *Services {
	userService := NewUserService(deps.Repositories.Users, deps.Hasher, deps.TokenManager, deps.TokenTTL, deps.UserCache)
	dialogService := NewDialogService(deps.Repositories.Dialogs)

	return &Services{
		UserService:   userService,
		DialogService: dialogService,
	}
}
