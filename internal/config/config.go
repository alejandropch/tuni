package config

import (
	"database/sql"
)

type AppConfig struct {
	DB *sql.DB
}

func New(db *sql.DB) *AppConfig {
	return &AppConfig{DB: db}
}
