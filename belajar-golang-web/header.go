package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			contentType := r.Header.Get("content-type")
			fmt.Fprintln(w, "Content-Type", contentType)

			contentLength := r.Header.Get("content-length")
			fmt.Fprintln(w, "Content-Length", contentLength)

			header := r.Header
			fmt.Fprintln(w, header)
		}),
	}
	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}