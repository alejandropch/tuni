package client

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	DB *sql.DB
}

func New(db *sql.DB) *Server {
	return &Server{
		DB: db,
	}
}

func Init() *Server {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	server := New(db)
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT
	);`

	_, err = server.DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("error creating table:\t%query: %s\n", err, sqlStmt)
	}
	return server
}
