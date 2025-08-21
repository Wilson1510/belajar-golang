package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Query().Get("name") != "" {
				http.ServeFile(w, r, "resources/ok.html")
			} else {
				http.ServeFile(w, r, "resources/notfound.html")	
			}
		}),
	}
	fmt.Println("Server is running on port 5300")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
