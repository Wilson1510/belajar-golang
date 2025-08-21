package main

import (
	"fmt"
	"net/http"

	"embed"
	"io/fs"
)

//go:embed resources
var sumber embed.FS

func main() {
	// directory := http.Dir("./resources")
	directory, _ := fs.Sub(sumber, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

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