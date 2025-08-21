package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl, err := template.ParseFiles("templates/name.html")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data := map[string]interface{}{
				"title": "Selamat datang di halaman template Go!",
				"name": "John Doe",
				"address": map[string]interface{}{
					"street": "Jl. Raya",
					"city": "Jakarta",
				},
			}

			// Eksekusi template dengan data string
			err := tmpl.ExecuteTemplate(w, "name.html", data)
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