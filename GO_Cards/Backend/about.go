package backend

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("Frontend/about.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
	}
}
