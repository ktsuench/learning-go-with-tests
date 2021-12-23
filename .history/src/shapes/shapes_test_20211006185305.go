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
		want := 62.8318530718
		checkPerimeter(t, want, circ)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, want float64, s Shape) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %g want %g given %v", got, want, s)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{10.0, 5.0}
		want := 50.0
		checkArea(t, want, rec)
	})
	t.Run("circles", func(t *testing.T) {
		circ := Circle{10}
		want := 314.1592653589793
		checkArea(t, want, circ)
	})
}
