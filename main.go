package main

import (
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, title FROM todos")
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
	initDB()
	defer DB.Close()

	http.HandleFunc("/", indexHandler)
	//	http.HandleFunc("/create", createHandler)
	//http.HandleFunc("/delete", deleteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
