package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 5.0}, hasPerimeter: 30.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasPerimeter: 36.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("got %g want %g given %#v", got, tt.hasPerimeter, tt.shape)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 5.0}, hasArea: 50.0},
		{shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got %g want %g given %#v", got, tt.hasArea, tt.shape)
		}
	}
}
