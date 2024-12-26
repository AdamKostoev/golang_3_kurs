package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// Инициализация базы данных
var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}

	// Создание таблиц
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		name TEXT NOT NULL,
		done INTEGER NOT NULL DEFAULT 0,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
