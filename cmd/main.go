package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
