package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	tmpl := template.New("global")
	tmpl = tmpl.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(name string) string {
			return strings.ToUpper(name)
		},
	})
	tmpl, _ = tmpl.Parse(`{{sayHello "john" | upper}}`)

	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := tmpl.ExecuteTemplate(w, "global", nil)
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