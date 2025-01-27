package backend

import (
	"html/template"
	"net/http"
	"time"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, short_description, published_at FROM news ORDER BY published_at DESC LIMIT 3")
	if err != nil {
		http.Error(w, "Ошибка при получении новостей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var recentNews []struct {
		ID               int
		Title            string
		ShortDescription string
		PublishedAt      time.Time
	}

	for rows.Next() {
		var n struct {
			ID               int
			Title            string
			ShortDescription string
			PublishedAt      time.Time
		}
		if err := rows.Scan(&n.ID, &n.Title, &n.ShortDescription, &n.PublishedAt); err != nil {
			http.Error(w, "Ошибка при чтении данных", http.StatusInternalServerError)
			return
		}
		recentNews = append(recentNews, n)
	}

	tmpl := template.Must(template.ParseFiles("Frontend/mainPage.html"))
	err = tmpl.Execute(w, recentNews)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
	}
}
