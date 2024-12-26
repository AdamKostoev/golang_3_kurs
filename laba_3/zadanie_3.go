package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Обработчик для POST-запроса
	r.POST("/count", func(c *gin.Context) {
		var requestBody struct {
			Text string `json:"text"`
		}

		// Прерывание в случае ошибки при чтении тела запроса
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный JSON"})
			return
		}

		// Создание карты для хранения количества вхождений каждого символа
		charCount := make(map[string]int)
		for _, char := range requestBody.Text {
			charCount[string(char)]++
		}

		// Возвращаем ответ в формате JSON
		c.JSON(http.StatusOK, charCount)
	})

	r.Run(":8083")
}
