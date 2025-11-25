package config

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	once      sync.Once
	appConfig *AppConfig
)

type AppConfig struct {
	DB *sql.DB
}

func New() *AppConfig {
	once.Do(func() {
		db, err := Init()
		if err != nil {
			panic(err)
		}
		appConfig = &AppConfig{
			DB: db,
		}
	})
	return appConfig
}

func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := db.Ping(); err == nil {
		log.Fatal(err)
		return nil, err
	}
	/*
		sqlStmt := `
		CREATE TABLE IF NOT EXISTS todos (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title TEXT
		);`

		_, err = server.DB.Exec(sqlStmt)
		if err != nil {
			log.Fatalf("error creating table:\t%query: %s\n", err, sqlStmt)
		}
	*/
	return db, nil
}
