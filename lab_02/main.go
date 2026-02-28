package main

import (
	"fmt"
	"math"
)

func arraysAndSlices() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := make([]int, 0, len(a))

	for i := range a {
		result = append(result, a[i]+b[i])
	}

	fmt.Println(result)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2.0
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
func (t Triangle) Perimeter() float64 {
	return 2 * (t.A + t.B + t.C)
}

func structuresAndInterfaces() {
	circle := Circle{
		Radius: 10,
	}
	rectangle := Rectangle{
		Width:  20,
		Height: 20,
	}
	triangle := Triangle{
		A: 20,
		B: 30,
		C: 30,
	}

	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
	fmt.Printf("Triangle Area: %.2f\n", triangle.Area())
}

func main() {
	arraysAndSlices()

	structuresAndInterfaces()
}
