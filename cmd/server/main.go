package main

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	assets "github.com/wham/minimalist-go-react-app/v2"
	"github.com/wham/minimalist-go-react-app/v2/internal/storage"
)

func main() {
	slog.Info("Starting server...")

	todos := storage.Todos{}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("HTML request", "method", r.Method, "path", r.RequestURI)

		template, err := template.ParseFS(assets.TemplatesFS, "templates/**.html")
		if err != nil {
			slog.Error("Failed to parse HTML templates", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}

		if err := template.ExecuteTemplate(w, "index.html", struct{}{}); err != nil {
			slog.Error("Failed to execute template", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
			return
		}
	})

	http.HandleFunc("GET /static/{name...}", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Static file request", "method", r.Method, "path", r.RequestURI)
		http.ServeFileFS(w, r, assets.StaticFS, "static/"+r.PathValue("name"))
	})

	http.HandleFunc("GET /ui.js", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("UI bundle request", "method", r.Method, "path", r.RequestURI)

		w.Write(assets.ReadUI())
	})

	http.HandleFunc("GET /api/todos", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(todos.List())
		w.Write(data)
	})

	http.HandleFunc("PUT /api/todos", func(w http.ResponseWriter, r *http.Request) {
		var todo storage.Todo
		json.NewDecoder(r.Body).Decode(&todo)
		todos.Add(todo)
	})

	slog.Error("Failed to start server", "error", http.ListenAndServe(":8080", nil))
	os.Exit(1)
}
