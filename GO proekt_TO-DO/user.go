package main

import (
	"golang.org/x/crypto/bcrypt"
)

// Регистрация нового пользователя
func registerUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	return err
}

// Проверка логина пользователя
func loginUser(username, password string) (int, error) {
	var storedPassword string
	var userID int
	err := db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&userID, &storedPassword)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		return 0, err
	}

	return userID, nil
}
