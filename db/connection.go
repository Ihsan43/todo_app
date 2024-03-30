package db

import (
	"database/sql"
	"log"

	_"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() *sql.DB {
	dsn := "user=postgres password=postgres host=localhost port=5432 dbname=todo_app sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed connect to database: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error when %s:", err.Error())
	}

	DB = db

	return db
}
