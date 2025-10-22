package main

import (
	"database/sql"
	"log"
)

type Todo struct {
	ID    int
	Title string
}

var DB *sql.DB

func initDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTENGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("error creating table:\t%query: %s\n", err, sqlStmt)
	}
}
