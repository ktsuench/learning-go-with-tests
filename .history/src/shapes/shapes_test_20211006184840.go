package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 5.0}
	got := rec.Perimeter()
	want := 30.0

	if got != want {
		t.Errorf("got %.1f want %.1f given %v", got, want, rec)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, want float64, s Shape) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %.1f want %1.f given %v", got, want, s)
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
