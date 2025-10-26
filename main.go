package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	client "tuni/database"
)

type Todo struct {
	ID    int
	Title string
}

var server *client.Server

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		_, err := server.DB.Exec("INSERT INTO todos(title) VALUES(?)", title)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := server.DB.Exec("DELETE FROM todos WHERE id = (?)", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := server.DB.Query("SELECT id, title FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	todos := []Todo{}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	tmpl := template.Must(template.New("index").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>Todo List</title>
	</head>
	<body>
	<h1>Todo List</h1>
	<form action="/create" method="POST">
	<input type="text" name="title" placeholder="New Todo" required>
	<button type="submit">Add</button>
	</form>
	<ul>
	{{range .}}
	<li>{{.Title}} <a href="/delete?id={{.ID}}">Delete</a></li>
	{{end}}
	</ul>
	</body>
	</html>
	`))
	tmpl.Execute(w, todos)
}
func main() {
	server = client.Init()
	defer server.DB.Close()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)
	fmt.Println("running bro")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
