package idler_service

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/model/entity"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/eugene-krivtsov/idler/internal/repository/redis"
	"github.com/eugene-krivtsov/idler/internal/server"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/eugene-krivtsov/idler/internal/transport/http"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	fmt.Println(`
 ================================================
 \\\   ######~~#####~~~##~~~~~~#####~~~#####   \\\
  \\\  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   \\\
   ))) ~~##~~~~##~~##~~##~~~~~~####~~~~#####     )))
  ///  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   ///
 ///   ######~~#####~~~######~~#####~~~##~~##  ///
 ================================================
	`)

	cfg, err := config.Init(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(config.PostgresConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		DB:       cfg.Postgres.DB,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	tokenManager := auth.NewJWTTokenManager(cfg.Auth.JWT.SigningKey)

	redisClient := redis.NewRedisClient(cfg.Redis)
	userCache := cache.NewRedisCache[string, entity.User](redisClient, cfg.Redis.Expires)
	//redis.NewUserCache(redisClient, cfg.Redis.Expires)

	repositories := repository.NewRepositories(db)
	services := service.NewServices(service.ServicesDepends{
		Repositories: repositories,
		Hasher:       hasher,
		TokenManager: tokenManager,
		TokenTTL:     cfg.Auth.JWT.TokenTTL,
		UserCache:    userCache,
	})
	handlers := http.NewHandler(services.UserService, tokenManager, services.DialogService)
	srv := server.NewServer(cfg, handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port))

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Print("IDLER facade application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("IDLER facade application shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	if err := redisClient.Close(); err != nil {
		logrus.Errorf("error occured on redic client close: %s", err.Error())
	}
}
