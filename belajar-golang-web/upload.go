package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"text/template"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// tampilkan form
		tmpl := template.Must(template.ParseFiles("templates/upload.html"))
		tmpl.ExecuteTemplate(w, "upload.html", nil)
		return
	}

	if r.Method == http.MethodPost {
		// ambil input name
		name := r.FormValue("name")

		// ambil file
		file, handler, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// simpan file ke folder "uploads/"
		dst, err := os.Create("resources/" + handler.Filename)
		if err != nil {
			panic(err)
		}
		defer dst.Close()

		// copy isi file ke tujuan
		_, err = io.Copy(dst, file)
		if err != nil {
			panic(err)
		}

		// tampilkan hasil sukses
		fmt.Fprintf(w, "Upload berhasil!<br>Nama: %s<br>File: %s", name, handler.Filename)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/form", Upload)

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