package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	assets "github.com/wham/minimalist-go-react-app/v2"
	"github.com/wham/minimalist-go-react-app/v2/internal/storage"
)

func main() {
	fmt.Println("Server is starting...")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(assets.TemplatesFS, "templates/**.html")

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		err = t.ExecuteTemplate(w, "index.html", struct{}{})

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
	})

	http.HandleFunc("GET /static/{name...}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, assets.StaticFS, r.PathValue("name"))
	})

	http.HandleFunc("GET /ui.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write(assets.ReadUI())
	})

	http.HandleFunc("GET /api/todos", func(w http.ResponseWriter, r *http.Request) {
		todos := storage.GetTodos()
		json, _ := json.Marshal(todos)
		w.Write(json)
	})

	http.HandleFunc("PUT /api/todos", func(w http.ResponseWriter, r *http.Request) {
		todo := storage.Todo{}
		json.NewDecoder(r.Body).Decode(&todo)
		storage.AddTodo(todo)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
