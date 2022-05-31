package facade_app

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/repository"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres"
	"github.com/eugene-krivtsov/idler/internal/server"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/eugene-krivtsov/idler/internal/transport/rest/handler"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	db, err := postgres.NewPostgresDB(postgres.Config{
		Port:     viper.GetString("postgres.port"),
		Username: viper.GetString("postgres.username"),
		DBName:   viper.GetString("postgres.dbname"),
		SSLMode:  viper.GetString("postgres.sslmode"),
		//Host:     os.Getenv("IP_ADDRESS"),
		Host: "localhost",
		//Password: os.Getenv("POSTGRES_PASSWORD"),
		Password: "qwerty",
	})

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	tokenManager := auth.NewJWTTokenManager(cfg.Auth.JWT.SigningKey)

	repositories := repository.NewRepositories(db)
	services := service.NewServices(service.ServicesDepends{
		Repositories: repositories,
		Hasher:       hasher,
		TokenManager: tokenManager,
		TokenTTL:     cfg.Auth.JWT.TokenTTL,
	})
	handlers := handler.NewHandler(services.Users)
	srv := server.NewServer(cfg, handlers.Init(cfg.HTTP.Host, cfg.HTTP.Port))

	go func() {
		if err := srv.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Print("IDLER application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("application shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
