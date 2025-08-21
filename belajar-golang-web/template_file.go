package main

import (
	"fmt"
	"html/template"
	"net/http"
	"embed"
)

//go:embed templates/*.gohtml
var templates embed.FS

func main() {
	// tmpl, err := template.ParseFiles("templates/simple.gohtml")
	// tmpl, err := template.ParseGlob("templates/*.gohtml")
	tmpl, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data := "Selamat datang di halaman template Go!"

			// Eksekusi template dengan data string
			err := tmpl.ExecuteTemplate(w, "simple.gohtml", data)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}),
	}

	fmt.Println("Server is running on port 5300")

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}