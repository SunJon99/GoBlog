package handle

import (
	"html/template"
	"net/http"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("templates/about.html"))

	if err := templates.ExecuteTemplate(w, "about.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
