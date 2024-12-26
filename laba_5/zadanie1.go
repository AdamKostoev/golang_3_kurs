package main

import (
	"fmt"
	"sync"
)

// Функция count, которая будет запускаться в горутине.
func count(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик после завершения работы горутины

	for num := range ch {

		fmt.Println("Square:", num*num) // - квадрат числа
	}
}

func main() {
	// Создаем канал для передачи чисел
	ch := make(chan int)

	var wg sync.WaitGroup

	// Запускаем горутину count
	wg.Add(1)
	go count(ch, &wg)

	// Отправляем несколько чисел в канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}

	// Закрываем канал, чтобы горутина могла завершить выполнение
	close(ch)

	// Ожидаем завершения горутины count
	wg.Wait()
}
