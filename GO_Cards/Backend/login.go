package backend

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	store = sessions.NewCookieStore([]byte("default-key"))
)

// LoginHandler обрабатывает страницу входа
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("Frontend/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		// Получаем данные из формы
		username := r.FormValue("username")
		password := r.FormValue("password")
		sessionDuration := r.FormValue("session_duration")

		var storedPassword string
		var role string

		// Получаем хеш пароля и роль пользователя
		query := "SELECT password, role FROM users WHERE username = $1"
		err := db.QueryRow(query, username).Scan(&storedPassword, &role)
		if err != nil {
			http.Error(w, "Ошибка аутентификации", http.StatusUnauthorized)
			return
		}

		// Сравниваем пароли
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
		if err != nil {
			http.Error(w, "Неверный пароль", http.StatusUnauthorized)
			return
		}

		// Создаем сессию
		session, _ := store.Get(r, "session")
		session.Values["username"] = username
		session.Values["role"] = role

		// Устанавливаем время жизни сессии
		var expiration time.Duration
		switch sessionDuration {
		case "10m":
			expiration = 10 * time.Minute
		case "1h":
			expiration = 1 * time.Hour
		case "5h":
			expiration = 5 * time.Hour
		case "24h":
			expiration = 24 * time.Hour
		default:
			expiration = 1 * time.Hour // По умолчанию - 1 час
		}

		session.Options.MaxAge = int(expiration.Seconds())
		session.Save(r, w)

		// Перенаправляем в зависимости от роли
		if role == "admin" {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		}
	}
}
