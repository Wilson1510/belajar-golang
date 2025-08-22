package main

import (
	"fmt"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("file") == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "File not found")
		return
	}
	w.Header().Add("Content-Disposition", "attachment; filename=tes123.jpg")
	http.ServeFile(w, r, "resources/gambar.png")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", Download)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("Server is running on port 8080")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}