package main

import (
	"fmt"
	"html/template"
	"embed"
	"net/http"
)

//go:embed templates/*.html
var templates embed.FS
var tmpl = template.Must(template.ParseFS(templates, "templates/*.html"))

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := tmpl.ExecuteTemplate(w, "escape.html", map[string]interface{}{
				"Title": "XSS",
				"Content": "<script>alert('Hello World')</script>",
				"Body": template.HTML("<h1>Hello World</h1>"),
			})
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