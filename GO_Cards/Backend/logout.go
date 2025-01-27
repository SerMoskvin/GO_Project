package backend

import (
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Ошибка при выходе", http.StatusInternalServerError)
		return
	}

	// Удаляем данные сессии
	session.Values["username"] = nil
	session.Values["role"] = nil
	session.Options.MaxAge = -1 // Удаляем сессию
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther) // Перенаправляем на страницу входа
}
