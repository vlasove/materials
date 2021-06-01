package geometry

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	allTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: &Rectangle{40.0, 20.0}, hasPerimeter: 120.0},
		{name: "Circle", shape: &Circle{40.0}, hasPerimeter: 2 * math.Pi * 40.0},
	}
	for _, tt := range allTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
		}
	}

}

func TestArea(t *testing.T) {
	allTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Rectangle",
			shape:   &Rectangle{10.0, 20.0},
			hasArea: 200.0,
		},
		{
			name:    "Circle",
			shape:   &Circle{40.0},
			hasArea: math.Pi * 40.0 * 40.0,
		},
	}
	for _, tt := range allTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
