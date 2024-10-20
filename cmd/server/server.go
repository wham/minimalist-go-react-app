package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wham/minimalist-go-react-app/v2/internal/assets"
	"github.com/wham/minimalist-go-react-app/v2/internal/storage"
)

func main() {
	fmt.Println("Server is starting...")

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		t, err := assets.ParseTemplate()

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
		content, err := assets.ReadStaticFile(r.PathValue("name"))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
		} else {
			w.Write(content)
		}
	})

	http.HandleFunc("GET /ui.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write(assets.ReadUiJs())
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
