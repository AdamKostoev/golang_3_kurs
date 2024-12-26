package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

func main() {
	// Открываем изображение
	inputFile, err := os.Open("laba_5/zadanie_4_vhodnoe_izobrazhenie.png")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	src := image.NewRGBA(bounds)
	draw.Draw(src, bounds, img, bounds.Min, draw.Src)

	// Создаём выходное изображение
	dest := image.NewRGBA(bounds)

	// Ядро свёртки Гаусса 3x3
	kernel := [][]float64{
		{0.0625, 0.125, 0.0625},
		{0.125, 0.25, 0.125},
		{0.0625, 0.125, 0.0625},
	}

	// Параллельная обработка
	applyConvolutionParallel(src, dest, kernel)

	// Сохраняем результат
	outputFile, err := os.Create("laba_5/res_zadanie_4.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, dest)
	if err != nil {
		panic(err)
	}
}

func applyConvolutionParallel(src, dest *image.RGBA, kernel [][]float64) {
	bounds := src.Bounds()
	var wg sync.WaitGroup

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				dest.Set(x, y, applyKernel(src, x, y, kernel))
			}
		}(y)
	}

	wg.Wait()
}

func applyKernel(img *image.RGBA, x, y int, kernel [][]float64) color.Color {
	kernelSize := len(kernel)
	offset := kernelSize / 2
	bounds := img.Bounds()

	var r, g, b, a float64

	for ky := 0; ky < kernelSize; ky++ {
		for kx := 0; kx < kernelSize; kx++ {
			ix := x + kx - offset
			iy := y + ky - offset

			// Проверяем, находится ли индекс в пределах изображения
			if ix >= bounds.Min.X && ix < bounds.Max.X && iy >= bounds.Min.Y && iy < bounds.Max.Y {
				pixel := img.RGBAAt(ix, iy)
				weight := kernel[ky][kx]

				r += float64(pixel.R) * weight
				g += float64(pixel.G) * weight
				b += float64(pixel.B) * weight
				a += float64(pixel.A) * weight
			}
		}
	}

	return color.RGBA{
		R: uint8(clamp(r)),
		G: uint8(clamp(g)),
		B: uint8(clamp(b)),
		A: uint8(clamp(a)),
	}
}

func clamp(value float64) float64 {
	if value > 255 {
		return 255
	}
	if value < 0 {
		return 0
	}
	return value
}
