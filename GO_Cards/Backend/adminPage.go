package backend

import (
	"html/template"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправляем на страницу входа, если не администратор
		return
	}

	rows, err := db.Query("SELECT id, title, short_description, published_at FROM news")
	if err != nil {
		http.Error(w, "Ошибка при получении новостей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var newsList []News
	for rows.Next() {
		var news News
		if err := rows.Scan(&news.ID, &news.Title, &news.Short_description, &news.Published_at); err != nil {
			http.Error(w, "Ошибка при чтении данных", http.StatusInternalServerError)
			return
		}
		newsList = append(newsList, news)
	}

	tmpl := template.Must(template.ParseFiles("Frontend/admin.html"))
	err = tmpl.Execute(w, newsList)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
		return
	}
}
