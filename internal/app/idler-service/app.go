package idler_service

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"github.com/eugene-krivtsov/idler/internal/model/dto"
	mongo_repository "github.com/eugene-krivtsov/idler/internal/repository/mongo-repository"
	"github.com/eugene-krivtsov/idler/internal/repository/postgres-repositrory"
	"github.com/eugene-krivtsov/idler/internal/repository/redis"
	"github.com/eugene-krivtsov/idler/internal/server"
	"github.com/eugene-krivtsov/idler/internal/service"
	"github.com/eugene-krivtsov/idler/internal/transport/http"
	"github.com/eugene-krivtsov/idler/internal/transport/ws"
	"github.com/eugene-krivtsov/idler/pkg/auth"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/eugene-krivtsov/idler/pkg/db/mongo"
	"github.com/eugene-krivtsov/idler/pkg/db/postgres"
	"github.com/eugene-krivtsov/idler/pkg/hash"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

// @title        Idler Application
// @version      1.0
// @description  API Server for keep in touch

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// Run initializes whole application

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

	postgresDB, err := postgres.NewPostgresDB(config.PostgresConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		DB:       cfg.Postgres.DB,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	mongoClient := mongo.NewMongoDb(cfg.Mongo)
	mongoDB := mongoClient.Database(cfg.Mongo.DB)

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	tokenManager := auth.NewJWTTokenManager(cfg.Auth.JWT.SigningKey)

	redisClient := redis.NewRedisClient(cfg.Redis)
	userCache := cache.NewRedisCache[string, dto.UserDTO](redisClient, cfg.Redis.Expires)

	postgresRepositories := postgres_repositrory.NewRepositories(postgresDB)
	mongoRepositories := mongo_repository.NewRepositories(mongoDB)

	services := service.NewServices(service.ServicesDepends{
		PostgresRepositories: postgresRepositories,
		MongoRepositories:    mongoRepositories,
		Hasher:               hasher,
		TokenManager:         tokenManager,
		TokenTTL:             cfg.Auth.JWT.TokenTTL,
		UserCache:            userCache,
	})
	httpHandlers := http.NewHandler(services.UserService, tokenManager, services.ConversationService)
	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(cfg.WS, hub, services.MessageService)
	wsServer := server.NewWSServer(cfg, wsHandler.Init(), hub)

	go func() {
		if err := wsServer.Run(); err != nil {
			logrus.Errorf("error occurred while running web socket server: %s\n", err.Error())
		}
	}()

	hub.Run()

	logrus.Print("IDLER facade application has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("IDLER facade application shutting down")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on http server shutting down: %s", err.Error())
	}

	if err := wsServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on ws server shutting down: %s", err.Error())
	}

	if err := postgresDB.Close(); err != nil {
		logrus.Errorf("error occured on postgres connection close: %s", err.Error())
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on mongo connection close: %s", err.Error())
	}

	if err := redisClient.Close(); err != nil {
		logrus.Errorf("error occured on redic client close: %s", err.Error())
	}
}
