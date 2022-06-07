package mongo

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	timeout = 10 * time.Second
)

func NewMongoDb(cfg config.MongoConfig) (*mongo.Client, error) {
	mongoUri := fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri).SetAuth(options.Credential{
		Username: cfg.User, Password: cfg.Password,
	}))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
