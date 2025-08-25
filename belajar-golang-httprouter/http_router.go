package main

import (
	"fmt"
	"net/http"
	"embed"
	"io/fs"
	"github.com/julienschmidt/httprouter"
)

//go:embed sumber
var resources embed.FS

func main() {
	router := httprouter.New()

	router.GET("/images/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintln(w, "Hello World")
		fmt.Fprintln(w, params)
	})

	router.GET("/products/*product", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintln(w, params)
	})

	dir, _ := fs.Sub(resources, "sumber")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	server := http.Server{
		Addr: "localhost:5300",
		Handler: router,
	}

	fmt.Println("Server running on :5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}