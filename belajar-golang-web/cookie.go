package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Powered-By"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"
	http.SetCookie(w, cookie)

	fmt.Fprintln(w, "Cookie created", cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-Powered-By")
	if err != nil {
		fmt.Fprintln(w, "Cookie not found")
	} else {
		name := cookie.Value
		fmt.Fprintln(w, "Cookie found", name)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", setCookie)
	mux.HandleFunc("/get-cookie", getCookie)

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