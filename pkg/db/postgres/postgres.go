package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	driverName = "postgres"
)

type Config struct {
	Host     string
	Port     string
	DB       string
	User     string
	Password string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(driverName, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DB, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
