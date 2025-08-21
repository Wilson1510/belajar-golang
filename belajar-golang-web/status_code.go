package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			name := query.Get("name")
			age := query.Get("age")

			if name == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Name is required")
			} else {
				fmt.Fprintln(w, "Hello ", name, "Your age is ", age)
			}
		}),
	}
	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}