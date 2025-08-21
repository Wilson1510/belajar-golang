package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}

			firstName := r.PostForm.Get("first_name")
			lastName := r.PostForm.Get("last_name")

			fmt.Fprintln(w, "Hello", firstName, lastName)
		}),
	}
	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}