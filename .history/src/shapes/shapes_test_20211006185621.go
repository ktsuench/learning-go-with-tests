package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, want float64, s Shape) {
		t.Helper()
		got := s.Perimeter()
		if got != want {
			t.Errorf("got %g want %g given %v", got, want, s)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{10.0, 5.0}
		want := 30.0
		checkPerimeter(t, want, rec)
	})
	t.Run("circles", func(t *testing.T) {
		circ := Circle{10.0}
		want := 62.83185307179586
		checkPerimeter(t, want, circ)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 5.0}, 50.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g given %v", got, tt.want, tt.shape)
		}
	}
}
