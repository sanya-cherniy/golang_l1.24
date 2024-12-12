package main

import (
	"fmt"
	"math"
)

// структура точки
type point struct {
	x float64
	y float64
}

// конструктор
func newPoint(x, y float64) *point {
	return &point{x: x, y: y}
}

// функция нахождения расстояния между точками
func (p *point) distance(other point) float64 {
	return math.Sqrt(math.Pow((other.x-p.x), 2) + math.Pow((other.y-p.y), 2))
}
func main() {
	point1 := newPoint(3.0, 2.0)
	point2 := newPoint(0.0, 0.0)
	fmt.Printf("distance = %f", point2.distance(*point1))
}
