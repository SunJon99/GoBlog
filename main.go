package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//静态文件服务
	files := http.FileServer(http.Dir("static")) //这里不需要使用 /static 详细说明可以看文档
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//这个是测试的  about页面
	mux.HandleFunc("/", index)
	mux.HandleFunc("/admin", admin)
	fmt.Println("Listening")
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	fmt.Println(server.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/about.html"))

	if err := templates.ExecuteTemplate(w, "about.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/admin/bolgs.html"))

	if err := templates.ExecuteTemplate(w, "bolgs.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
