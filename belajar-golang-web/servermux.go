package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About Page")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Images Page")
	})

	mux.HandleFunc("/images/thumbnail/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thumbnail Page")
	})


	server := http.Server{
		Addr: "localhost:5300",
		Handler: mux,
	}

	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}