package geometry

import "math"

// Rectangle ...
type Rectangle struct {
	Width  float64
	Length float64
}

// Perimeter ...
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Length)
}

// Area ...
func (r *Rectangle) Area() float64 {
	return r.Width * r.Length
}

// Circle ...
type Circle struct {
	Radius float64
}

// Perimeter ...
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area ...
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Shape ...
type Shape interface {
	Area() float64
	Perimeter() float64
}
