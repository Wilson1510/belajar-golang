package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:5300",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, r.Method)
			fmt.Fprintln(w, r.URL.Path)
			fmt.Fprintln(w, r.Proto)
			fmt.Fprintln(w, r.Host)
			fmt.Fprintln(w, r.RemoteAddr)
			fmt.Fprintln(w, r.UserAgent())
			fmt.Fprintln(w, r.Referer())
		}),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}