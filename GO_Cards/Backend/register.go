package backend

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Отображаем страницу регистрации
		tmpl := template.Must(template.ParseFiles("Frontend/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		// Обработка данных формы
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Хеширование пароля
		hashedPassword, err := HashPassword(password)
		if err != nil {
			http.Error(w, "Ошибка хеширования пароля", http.StatusInternalServerError)
			return
		}

		user := User{
			Username: username,
			Password: hashedPassword,
			Role:     "user",
		}

		// Сохранение пользователя в базу данных
		err = SaveUserToDB(user)
		if err != nil {
			http.Error(w, "Ошибка сохранения пользователя", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// HashPassword хеширует пароль
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// SaveUser ToDB сохраняет пользователя в базе данных
func SaveUserToDB(user User) error {
	_, err := db.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", user.Username, user.Password, user.Role)
	return err
}
