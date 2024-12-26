package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

func main() {
	// Открываем изображение
	inputFile, err := os.Open("laba_5/zadanie_2_3_vhodnoe_izobrazhenie.png")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	// Преобразуем в RGBA64 для возможности изменения
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA64(bounds)
	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)

	// Последовательная обработка
	start := time.Now()
	filterSequential(rgbaImg)
	sequentialDuration := time.Since(start)
	saveImage("zadanie.3.output_sequential.png", rgbaImg)

	// Параллельная обработка
	start = time.Now()
	filterParallel(rgbaImg)
	parallelDuration := time.Since(start)
	saveImage("zadanie.3.output_parallel.png", rgbaImg)

	// Сравнение времени
	println("Sequential processing time:", sequentialDuration.String())
	println("Parallel processing time:", parallelDuration.String())
}

// Фильтр для обработки строки пикселей
func filterSequential(img *image.RGBA64) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			applyFilter(img, x, y)
		}
	}
}

// Функция параллельной обработки
func filterParallel(img *image.RGBA64) {
	var wg sync.WaitGroup
	bounds := img.Bounds()
	wg.Add(bounds.Max.Y - bounds.Min.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		go func(y int) {
			defer wg.Done()
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				applyFilter(img, x, y)
			}
		}(y)
	}

	wg.Wait()
}

// Применение фильтра к одному пикселю
func applyFilter(img *image.RGBA64, x, y int) {
	original := img.At(x, y).(color.RGBA64)
	newColor := color.RGBA64{
		R: 0xFFFF - original.R,
		G: 0xFFFF - original.G,
		B: 0xFFFF - original.B,
		A: original.A,
	}
	img.Set(x, y, newColor)
}

// Сохранение изображения
func saveImage(filename string, img *image.RGBA64) {
	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, img)
	if err != nil {
		panic(err)
	}
}
