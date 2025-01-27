package backend

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, title, short_description, published_at FROM news ORDER BY published_at DESC")
	if err != nil {
		http.Error(w, "Ошибка при получении новостей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allNews []struct {
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
		allNews = append(allNews, n)
	}

	tmpl := template.Must(template.ParseFiles("Frontend/news.html"))
	err = tmpl.Execute(w, allNews)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
	}
}

// Обрабатывает запросы для страницы с полным текстом новости
func NewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение ID из URL
	idStr := strings.TrimPrefix(r.URL.Path, "/news/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID новости", http.StatusBadRequest)
		return
	}

	// Запрос к базе данных для получения полной новости по ID
	var news struct {
		ID          int
		Title       string
		Content     string
		PublishedAt time.Time
	}

	err = db.QueryRow("SELECT id, title, content, published_at FROM news WHERE id = $1", id).Scan(&news.ID, &news.Title, &news.Content, &news.PublishedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Новость не найдена", http.StatusNotFound)
		} else {
			http.Error(w, "Ошибка при получении новости", http.StatusInternalServerError)
		}
		return
	}

	// Отображение шаблона с полной новостью
	tmpl := template.Must(template.ParseFiles("Frontend/newsDetail.html"))
	err = tmpl.Execute(w, news)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
	}
}
