package main

import (
	"database/sql"
	"tuni/internal/config"
)

type UniversityRepo struct {
	db *sql.DB
}

func New(db *sql.DB) {
	return &UniversityRepo{
		db: config.New().DB,
	}

}
