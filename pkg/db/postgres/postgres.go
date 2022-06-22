package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

const (
	driverName = "pgx"
)

type Config struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	SSLMode  string
}

func NewPostgresPool(cfg Config) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, err
	}

	config.ConnConfig.Host = cfg.Host
	config.ConnConfig.Port = 5432
	config.ConnConfig.Database = cfg.DB
	config.ConnConfig.User = cfg.User
	config.ConnConfig.Password = cfg.Password

	config.MaxConns = 4
	config.MinConns = 2
	config.MaxConnLifetime = 5 * time.Second
	config.MaxConnIdleTime = 5 * time.Second

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	return pool, nil
}
