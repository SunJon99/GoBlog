package main

import (
	"GoBlog/boxs"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	fmt.Println("Listening")
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/about.html"))

	if err := templates.ExecuteTemplate(w, "about.html", boxs.Blog{}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
