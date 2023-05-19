// Package repository has repository for application.
package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Config has configuration for connection to postgresql.
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB constructor for connection with configuration.
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode),
	)
	if err != nil {
		return nil, err
	}

	errPing := db.Ping()
	if err != nil {
		return nil, errPing
	}

	return db, nil
}
