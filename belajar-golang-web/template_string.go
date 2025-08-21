package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// Template string sederhana.
	// Placeholder {{.}} akan digantikan dengan data yang kita berikan.
	// Titik (.) digunakan saat data yang disuntikkan adalah tipe data dasar seperti string,
	// bukan struct atau map.
	const tmplStr = `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Contoh Tanpa HTML File</title>
		</head>
		<body>
			<h1>{{.}}</h1>
		</body>
		</html>`

	// Parsing template dari string
	tmpl, err := template.New("webpage").Parse(tmplStr)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data := "Selamat datang di halaman template Go!"

			// Eksekusi template dengan data string
			err := tmpl.ExecuteTemplate(w, "webpage", data)
			if err != nil {
				fmt.Printf("Error executing template: %v\n", err)
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