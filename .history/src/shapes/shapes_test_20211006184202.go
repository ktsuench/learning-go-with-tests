package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 5.0}
	got := Perimeter(rec)
	want := 30.0

	if got != want {
		t.Errorf("got %.1f want %.1f given %v", got, want, rec)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{10.0, 5.0}
		got := Area(rec)
		want := 50.0

		if got != want {
			t.Errorf("got %.1f want %1.f given %v", got, want, rec)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circ := Circle{10}
		got := Area(circ)
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g given %v", got, want, circ)
		}
	})
}
