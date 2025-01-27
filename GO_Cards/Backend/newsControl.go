package backend

import (
	"html/template"
	"net/http"
	"time"
)

// AddNewsHandler обрабатывает запросы на добавление новостей
func AddNewsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправляем на страницу входа, если не администратор
		return
	}

	if r.Method == http.MethodGet {
		// Если метод GET, отобразить форму для добавления новости
		tmpl := template.Must(template.ParseFiles("Frontend/add.html"))
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		// Если метод POST, сохранить новость в базе данных
		title := r.FormValue("title")
		shortDescription := r.FormValue("short_description")
		content := r.FormValue("content")
		publishedAt := time.Now()

		_, err := db.Exec("INSERT INTO news (title, short_description, content, published_at) VALUES ($1, $2, $3, $4)", title, shortDescription, content, publishedAt)
		if err != nil {
			http.Error(w, "Ошибка при добавлении новости", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther) // Перенаправляем обратно на админ-страницу
	}
}

func EditNewsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправляем на страницу входа, если не администратор
		return
	}

	newsID := r.URL.Path[len("/edit/"):]

	var news News
	err = db.QueryRow("SELECT id, title, short_description, content, published_at FROM news WHERE id = $1", newsID).Scan(&news.ID, &news.Title, &news.Short_description, &news.Content, &news.Published_at)
	if err != nil {
		http.Error(w, "Ошибка при получении новости", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("Frontend/edit.html"))
	err = tmpl.Execute(w, news)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
		return
	}
}

// UpdateNewsHandler обрабатывает обновление новости
func UpdateNewsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправляем на страницу входа, если не администратор
		return
	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		title := r.FormValue("title")
		shortDescription := r.FormValue("short_description")
		content := r.FormValue("content")

		_, err := db.Exec("UPDATE news SET title = $1, short_description = $2, content = $3 WHERE id = $4", title, shortDescription, content, id)
		if err != nil {
			http.Error(w, "Ошибка при обновлении новости", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther) // Перенаправляем обратно на админ-страницу
	}
}

// DeleteNewsHandler обрабатывает удаление новости по ID
func DeleteNewsHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["role"] != "admin" {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправляем на страницу входа, если не администратор
		return
	}

	newsID := r.URL.Path[len("/delete/"):]

	_, err = db.Exec("DELETE FROM news WHERE id = $1", newsID)
	if err != nil {
		http.Error(w, "Ошибка при удалении новости", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther) // Перенаправляем обратно на админ-страницу
}
