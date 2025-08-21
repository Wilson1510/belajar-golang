package main

import (
	"fmt"
	"html/template"
	"embed"
	"net/http"
)

//go:embed templates/*.gohtml
var templates embed.FS
var tmpl = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := tmpl.ExecuteTemplate(w, "simple.gohtml", "Hello World")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}),
	}

	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}