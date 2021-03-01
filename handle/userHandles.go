package handle

import (
	"html/template"
	"net/http"
)

// IndexHandle 用来处理用户的首页请求
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/about.html"))

	if err := templates.ExecuteTemplate(w, "about.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
