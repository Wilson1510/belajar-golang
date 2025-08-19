package main

import (
	"fmt"
	"net/http"
)

func main() {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:5300",
		Handler: handler,
	}

	fmt.Println("Server is running on port 5300")
	err := server.ListenAndServe()
	fmt.Println("Server is already running")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is shutdown")
}