package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

// 2. Открытие и декодирование изображения
func main() {
	// Открываем файл с изображением
	file, err := os.Open("laba_5/zadanie_2_3_vhodnoe_izobrazhenie.png")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Декодируем изображение
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка декодирования изображения:", err)
		return
	}

	// Преобразуем изображение в draw.RGBA64Image
	drawImg, ok := img.(draw.RGBA64Image)
	if !ok {
		fmt.Println("Ошибка преобразования изображения")
		return
	}

	// Замер времени до начала обработки
	start := time.Now()

	// Применяем фильтр (параллельно)
	filter(drawImg)

	// Замер времени после обработки
	elapsed := time.Since(start)
	fmt.Printf("Время обработки: %s\n", elapsed)

	// Сохраняем результат
	outputFile, err := os.Create("res_zadanie.2.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outputFile.Close()

	// Кодируем изображение в файл
	err = png.Encode(outputFile, drawImg)
	if err != nil {
		fmt.Println("Ошибка кодирования изображения:", err)
		return
	}

	fmt.Println("Обработанное изображение сохранено в res_zadanie.2.png")
}

// 3. Функция фильтрации (перевод в оттенки серого)
// filter применяет преобразование к каждому пикселю изображения
func filter(img draw.RGBA64Image) {
	bounds := img.Bounds()
	var wg sync.WaitGroup

	// Параллельная обработка
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				// Получаем текущий цвет пикселя
				c := img.RGBA64At(x, y)

				// Преобразуем в оттенки серого
				grayValue := uint16((int(c.R) + int(c.G) + int(c.B)) / 3)

				// Создаем новый цвет и записываем его
				grayColor := color.RGBA64{
					R: grayValue,
					G: grayValue,
					B: grayValue,
					A: c.A, // альфа-канал остаётся прежним
				}
				img.SetRGBA64(x, y, grayColor)
			}
		}(y)
	}

	// Ждем завершения всех горутин
	wg.Wait()
}
