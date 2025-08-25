package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprintln(w, "Terjadi panic", error)
	}

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halaman tidak ditemukan")
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Metode tidak diizinkan")
	})

	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	router.POST("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprintln(w, "POST")
	})

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