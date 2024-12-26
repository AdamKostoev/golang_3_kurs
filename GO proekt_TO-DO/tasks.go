package main

// Добавление задачи в базу данных
func addTask(userID int, taskName string) error {
	_, err := db.Exec("INSERT INTO tasks (user_id, name) VALUES (?, ?)", userID, taskName)
	return err
}

// Получение задач для пользователя
func getTasks(userID int) ([]Task, error) {
	rows, err := db.Query("SELECT id, name, done FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Обновление статуса задачи
func toggleTaskStatus(taskID int, newStatus bool) error {
	status := 0
	if newStatus {
		status = 1
	}
	_, err := db.Exec("UPDATE tasks SET done = ? WHERE id = ?", status, taskID)
	return err
}

// Удаление задачи
func deleteTask(taskID int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", taskID)
	return err
}
