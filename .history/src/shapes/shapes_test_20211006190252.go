package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 5.0}, 30.0},
		{Circle{10.0}, 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g given %v", got, tt.want, tt.shape)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 5.0}, 50.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g given %v", got, tt.want, tt.shape)
		}
	}
}
