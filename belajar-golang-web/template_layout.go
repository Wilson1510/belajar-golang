package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Title string
	Name string
}

func main() {
	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/header.html",
		"templates/footer.html",
	)
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data := User{
				Title: "Selamat datang di halaman template Go Struct!",
				Name: "John Doe",
			}

			// Eksekusi template dengan data string
			err := tmpl.ExecuteTemplate(w, "layout.html", data)
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