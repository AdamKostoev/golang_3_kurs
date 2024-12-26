package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Маршрут для сложения
	r.GET("/add", func(c *gin.Context) {
		a, b, err := getNumbers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := a + b
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// Маршрут для вычитания
	r.GET("/sub", func(c *gin.Context) {
		a, b, err := getNumbers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := a - b
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// Маршрут для умножения
	r.GET("/mul", func(c *gin.Context) {
		a, b, err := getNumbers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := a * b
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// Маршрут для деления
	r.GET("/div", func(c *gin.Context) {
		a, b, err := getNumbers(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if b == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Деление на 0 невозможно"})
			return
		}
		result := a / b
		c.JSON(http.StatusOK, gin.H{"result": result})
	})

	// Запускаем сервер на порту 8081
	r.Run(":8081")
}

// Функция для получения чисел из query-параметров
func getNumbers(c *gin.Context) (float64, float64, error) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	if aStr == "" || bStr == "" {
		return 0, 0, fmt.Errorf("необходимы оба параметра: a и b")
	}

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("параметр a должен быть числом")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("параметр b должен быть числом")
	}

	return a, b, nil
}
