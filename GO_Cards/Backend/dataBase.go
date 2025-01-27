package backend

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func InitDB() {
	var err error
	connStr := "user=postgres dbname=GO_Cards password=test host=localhost sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Ошибка при подключении к базе данных:", err)
	}

	CreateTables()

}

// Создание таблиц, если они не существуют
func CreateTables() {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(100) NOT NULL,
		role VARCHAR(50) NOT NULL
	);`

	newsTable := `
	CREATE TABLE IF NOT EXISTS news (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		short_description TEXT NOT NULL,
		content TEXT NOT NULL,
		published_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	cardsTable := `
	CREATE TABLE IF NOT EXISTS cards (
		id SERIAL PRIMARY KEY,
		card_name VARCHAR(100) NOT NULL,
		artistic_description TEXT NOT NULL,
		release VARCHAR(50) NOT NULL,
		card_number INT NOT NULL,
		card_type VARCHAR(50) NOT NULL,
		cost DECIMAL(10, 2) NOT NULL,
		element VARCHAR(50) NOT NULL,
		attack_power INT NOT NULL,
		health_points INT NOT NULL,
		card_description TEXT NOT NULL
	);`

	// Выполнение SQL-запросов для создания таблиц
	_, err := db.Exec(usersTable)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы пользователей: %v", err)
	}

	_, err = db.Exec(newsTable)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы новостей: %v", err)
	}

	_, err = db.Exec(cardsTable)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы карт: %v", err)
	}

}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Printf("Ошибка при закрытии базы данных: %v", err)
	}
}
