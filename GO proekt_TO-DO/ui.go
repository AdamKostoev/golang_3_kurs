package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Отображение задач
func showTasks(w fyne.Window, userID int) {
	tasks, err := getTasks(userID)
	if err != nil {
		log.Println("Error fetching tasks:", err)
		w.SetContent(widget.NewLabel("Ошибка получения задач"))
		return
	}

	// Виджеты для отображения задач
	var taskWidgets []fyne.CanvasObject
	for _, task := range tasks {
		taskCopy := task

		check := widget.NewCheck(taskCopy.Name, func(checked bool) {
			err := toggleTaskStatus(taskCopy.ID, checked)
			if err != nil {
				log.Println("Error updating task:", err)
			}
		})

		check.SetChecked(taskCopy.Done)
		taskWidgets = append(taskWidgets, check)

		editButton := widget.NewButton("Edit", func() {
			showEditTaskForm(w, taskCopy.ID, userID)
		})

		deleteButton := widget.NewButton("Delete", func() {
			err := deleteTask(taskCopy.ID)
			if err != nil {
				log.Println("Error deleting task:", err)
			} else {
				showTasks(w, userID)
			}
		})

		taskWidgets = append(taskWidgets, container.NewHBox(editButton, deleteButton))
	}

	taskEntry := widget.NewEntry()
	taskEntry.SetPlaceHolder("Enter new task")

	addTaskButton := widget.NewButton("Add Task", func() {
		taskName := taskEntry.Text
		if taskName != "" {
			err := addTask(userID, taskName)
			if err != nil {
				log.Println("Error adding task:", err)
			} else {
				taskEntry.SetText("")
				showTasks(w, userID)
			}
		}
	})

	logoutButton := widget.NewButton("Logout", func() {
		showLoginForm(w)
	})

	content := container.NewVBox(
		widget.NewLabel(fmt.Sprintf("User ID: %d", userID)),
		container.NewVBox(taskWidgets...),
		taskEntry,
		addTaskButton,
		logoutButton,
	)

	w.SetContent(content)
}

// Создание формы редактирования задачи
func showEditTaskForm(w fyne.Window, taskID, userID int) {
	taskNameEntry := widget.NewEntry()

	var taskName string
	err := db.QueryRow("SELECT name FROM tasks WHERE id = ?", taskID).Scan(&taskName)
	if err != nil {
		log.Println("Error fetching task:", err)
		return
	}
	taskNameEntry.SetText(taskName)

	saveButton := widget.NewButton("Save", func() {
		newTaskName := taskNameEntry.Text
		if newTaskName != "" {
			_, err := db.Exec("UPDATE tasks SET name = ? WHERE id = ?", newTaskName, taskID)
			if err != nil {
				log.Println("Error updating task:", err)
			} else {
				showTasks(w, userID)
			}
		}
	})

	cancelButton := widget.NewButton("Cancel", func() {
		showTasks(w, userID)
	})

	content := container.NewVBox(
		widget.NewLabel("Edit Task"),
		taskNameEntry,
		saveButton,
		cancelButton,
	)
	w.SetContent(content)
}

// Создание формы регистрации
func createRegistrationForm(w fyne.Window) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	registerButton := widget.NewButton("Register", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		err := registerUser(username, password)
		if err != nil {
			log.Println("Registration error:", err)
		} else {
			log.Println("User registered successfully!")
			showLoginForm(w)
		}
	})

	content := container.NewVBox(usernameEntry, passwordEntry, registerButton)
	w.SetContent(content)
}

// Создание формы логина
func createLoginForm(w fyne.Window) {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Username")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Password")

	loginButton := widget.NewButton("Login", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text
		userID, err := loginUser(username, password)
		if err != nil {
			log.Println("Login error:", err)
		} else {
			log.Println("Login successful!")
			showTasks(w, userID)
		}
	})

	registerButton := widget.NewButton("Register", func() {
		showRegistrationForm(w)
	})

	content := container.NewVBox(usernameEntry, passwordEntry, loginButton, registerButton)
	w.SetContent(content)
}

// Создание формы для авторизации и регистрации
func showLoginForm(w fyne.Window) {
	createLoginForm(w)
}

// Создание формы для регистрации
func showRegistrationForm(w fyne.Window) {
	createRegistrationForm(w)
}
