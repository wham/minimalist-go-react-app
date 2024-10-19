package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	fmt.Println("Server is starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/index.html")

		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}

		t.Execute(w, struct{}{})
	})

	http.HandleFunc("GET /static/{path...}", func(w http.ResponseWriter, r *http.Request) {
		path := "static/" + r.PathValue("path")
		content, err := os.ReadFile(path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
		} else {
			w.Write(content)
		}
	})

	http.HandleFunc("/ui.js", func(w http.ResponseWriter, r *http.Request) {
		result := api.Build(api.BuildOptions{
			EntryPoints: []string{"ui/main.tsx"},
			Bundle:      true,
			Format:      api.FormatESModule,
			Sourcemap:   api.SourceMapInline,
		})

		w.Write(result.OutputFiles[0].Contents)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
