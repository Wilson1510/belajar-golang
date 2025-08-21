package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Name string
}

func (page Page) SayHello(name string) string {
	return "Hello " + name + " from " + page.Name
}

func main() {
	tmpl, err := template.New("function").Parse(`{{.SayHello "John"}} {{len .Name}}`)
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			data := Page{
				Title: "Template Function",
				Name: "John Doe",
			}

			err := tmpl.ExecuteTemplate(w, "function", data)
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