package backend

import (
	"html/template"
	"net/http"
)

// HomeHandler обрабатывает домашнюю страницу пользователя
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil || session.Values["username"] == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther) // Если сессия не найдена, перенаправляем на страницу входа
		return
	}

	// Извлекаем имя пользователя из сессии
	username := session.Values["username"].(string)

	// Создаем структуру для передачи данных в шаблон
	data := struct {
		Username string
	}{
		Username: username,
	}

	// Загружаем HTML-шаблон для домашней страницы
	tmpl := template.Must(template.ParseFiles("Frontend/homePage.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка при загрузке страницы", http.StatusInternalServerError)
		return
	}
}
