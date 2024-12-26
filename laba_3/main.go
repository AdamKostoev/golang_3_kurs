package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер
	r := gin.Default()

	// Обрабатываем GET-запрос на корневом маршруте
	r.GET("/", func(c *gin.Context) {
		// Получаем query-параметры
		name := c.Query("name")
		age := c.Query("age")

		// Формируем ответ
		response := fmt.Sprintf("Меня зовут %s, мне %s лет", name, age)

		// Отправляем ответ клиенту
		c.String(200, response)
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
