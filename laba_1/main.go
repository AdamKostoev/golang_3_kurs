package main

import "fmt"

// Задание № 2
func hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

//Задание № 3
func printEven(n, n1 int64) error {
	if n > n1 {
		return fmt.Errorf("Ошибка:Левая граница диапазона больше правой")
	} else {
		for i := n; i <= n1; i++ {
			if i%2 == 0 {
				_, err := fmt.Println(i)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Задание № 4
func apply(a, b float64, operator string) (float64, error) {
	if operator == "+" {
		return a + b, nil
	}
	if operator == "-" {
		return a - b, nil
	}
	if operator == "*" {
		return a * b, nil
	}
	if operator == "/" {
		if b != 0 {
			return a / b, nil
		}
	}
	return 0, fmt.Errorf(", действие не поддерживается")
}

func main() {
	fmt.Println("Задание № 4")
	fmt.Println(apply(2, 3, "+"))
	fmt.Println(apply(2, 3, "/"))
	fmt.Println(apply(7, 10, "*"))
	fmt.Println(apply(2, 3, "#"))
	fmt.Println("Задание № 3")
	err := printEven(1, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Задание № 2")
	fmt.Println(hello("Адам"))
	fmt.Println("Задане № 1")
	fmt.Println("HEllo, World")
}
