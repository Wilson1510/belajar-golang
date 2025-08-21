package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			name := query["name"]
			age := query.Get("age")

			fmt.Fprintln(w, "Hello ", strings.Join(name, " "))
			fmt.Fprintln(w, "Your age is ", age)
		}),
	}
	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}