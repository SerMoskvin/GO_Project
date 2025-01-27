package backend

import (
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/admin", AdminHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/add", AddNewsHandler)
	http.HandleFunc("/edit/", EditNewsHandler)
	http.HandleFunc("/update", UpdateNewsHandler)
	http.HandleFunc("/delete/", DeleteNewsHandler)
	http.HandleFunc("/news", NewsHandler)
	http.HandleFunc("/news/", NewsDetailHandler)
	http.HandleFunc("/about", AboutHandler)

	http.ListenAndServe(":8080", nil)
}
