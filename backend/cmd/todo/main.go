package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/trite8q1/todo/backend/pkg/models"
)

var tmpl *template.Template

func todo(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title: "Todo List",
		Todos: []models.Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Learned about GORM", Done: false},
		},
	}

	tmpl.Execute(w, data)
}

func handleRoutes(router *http.ServeMux) {
	router.HandleFunc("/todo", todo)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("../frontend/static/templates/index.gohtml"))

	fs := http.FileServer(http.Dir("../frontend/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	handleRoutes(mux)

	log.Fatal(http.ListenAndServe(":9091", mux))
}
