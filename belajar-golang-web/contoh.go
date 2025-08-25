package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler biasa
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()

	// tambahkan route
	mux.HandleFunc("/", helloHandler)

	// bungkus mux dengan middleware
	loggedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// jalankan handler berikutnya
		mux.ServeHTTP(w, r)

		// log setelah request selesai
		fmt.Printf("%s %s %s\n", r.Method, r.URL.Path, time.Since(start))
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", loggedMux)
}
