package main

import (
	"errors"
	"fmt"
	"math"
)

// Задание № 1.1
func IPformat(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// Задание № 1.2
func listEven(left, right int) ([]int, error) {
	if left > right {
		return nil, errors.New("Левая граница больше правой")
	}

	var evenNumbers []int
	for i := left; i <= right; i++ {
		if i%2 == 0 {
			evenNumbers = append(evenNumbers, i)
		}
	}
	return evenNumbers, nil
}

// Задание № 2
func countChars(str string) map[rune]int {
	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}
	return charCount
}

// Задание № 3.1
type Point struct {
	X float64
	Y float64
}

// Задание № 3.2
type Segment struct {
	start Point
	end   Point
}

type SegmentArray struct {
	Points [2]Point
}

// Задание № 3.3
func (s Segment) Length() float64 {
	return math.Sqrt(math.Pow(s.end.X-s.start.X, 2) + math.Pow(s.end.Y-s.start.Y, 2))
}

// Задание № 3.4
type Triangle struct {
	Base   float64
	Height float64
}

// Задание № 3.5
type Circle struct {
	Radius float64
}

// Задание № 3.6
func (t Triangle) Area() float64 {
	return 05 * t.Base * t.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Задание № 3.7
type Shape interface {
	Area() float64
}

// Задание 3.8
func printArea(s Shape) {
	result := s.Area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

func main() {
	fmt.Println("Задание 1.1")
	ip := [4]byte{127, 0, 0, 1}
	fmt.Println(IPformat(ip))
	fmt.Println("Задание 1.2")
	fmt.Println(listEven(33, 23))
	fmt.Println("Задание 2")
	str := "Количество символов"
	fmt.Println(countChars(str))
	fmt.Println("Задание 3.1")
	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println("Задание 3.2")
	segment1 := Segment{
		start: Point{X: 1, Y: 2},
		end:   Point{X: 5, Y: 6},
	}
	segment2 := SegmentArray{
		Points: [2]Point{{X: 2, Y: 8}, {X: 3, Y: 7}},
	}
	fmt.Printf("Segment 1: Start(%v,%v), End(%v,%v)\n", segment1.start.X, segment1.start.Y, segment1.end.X, segment1.end.Y)
	fmt.Printf("Segment 2: Start(%v,%v), End(%v,%v)\n", segment2.Points[0].X, segment2.Points[0].Y, segment2.Points[1].X, segment2.Points[1].Y)
	fmt.Println("Задание 3.3")
	fmt.Println(segment1.Length())
	fmt.Println("Задание 3.4")
	fmt.Println("Задание 3.5")
	fmt.Println("Задание 3.6")
	triangle := Triangle{Base: 4, Height: 5}
	circle := Circle{Radius: 5}

	triangleArea := triangle.Area()
	circleArea := circle.Area()
	fmt.Println("Площадь треугольника: ", triangleArea)
	fmt.Println("Площадь круга: ", circleArea)
	fmt.Println("Задание 3.7")
	fmt.Println("Задание 3.8")
	printArea(triangle)
	printArea(circle)
}
