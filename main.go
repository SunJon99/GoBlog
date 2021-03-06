package main

import (
	handle "GoBlog/handle"
	"fmt"
	"net/http"
)

type test interface {
	test() string
	fuck()
}

type date struct{}

func (date) test() string  { return "" }
func (*date) test() string { return "" }
func (date) fuck()         {}
func main() {
	var d date
	var t test = d

	mux := http.NewServeMux()

	//静态文件服务
	static := http.FileServer(http.Dir("static")) //这里不需要使用 /static 详细说明可以看文档
	mux.Handle("/static/", http.StripPrefix("/static/", static))

	//这个是测试的  about页面
	mux.HandleFunc("/", handle.IndexHandle)
	mux.HandleFunc("/admin", handle.AdminHandle)
	mux.HandleFunc("/admin/edit", handle.EditHandle)
	mux.HandleFunc("/info/blogs", handle.GetBlogsHandle)
	fmt.Println("Listening")
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	fmt.Println(server.ListenAndServe())
}
