// Package task31 Написать программу нахождения расстояния между 2 точками,
// которые представление в виде структуры Point с
// инкапсулированными параметрами x,y и конструктором.
package task31

import (
	"fmt"
	"math"
)

// Point ...
type Point struct {
	x, y float64
}

// NewPoint ...
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// GetX ...
func (p *Point) GetX() float64 {
	return p.x
}

// GetY ...
func (p *Point) GetY() float64 {
	return p.y
}

// Distance ...
func Distance(p1, p2 *Point) float64 {
	deltaX := p1.GetX() - p2.GetX()
	deltaY := p1.GetY() - p2.GetY()

	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

// Start ...
func Start() {
	d := Distance(NewPoint(0, 0), NewPoint(10, 10))
	fmt.Println("Distance is:", d)
}
