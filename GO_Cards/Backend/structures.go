package backend

import "time"

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type News struct {
	ID                int
	Title             string
	Short_description string
	Content           string
	Published_at      time.Time
}
