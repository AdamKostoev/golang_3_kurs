package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	// Инициализация базы данных
	initDB()

	// Создание нового приложения
	a := app.New()
	w := a.NewWindow("To-Do App")

	// Отображаем форму логина
	showLoginForm(w)

	// Запуск приложения
	w.ShowAndRun()
}
