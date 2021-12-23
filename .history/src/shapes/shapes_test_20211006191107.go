package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 5.0}, want: 30.0},
		{shape: Circle{Radius: 10.0}, want: 62.83185307179586},
		{shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g given %#v", got, tt.want, tt.shape)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 5.0}, want: 50.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g given %#v", got, tt.want, tt.shape)
		}
	}
}
