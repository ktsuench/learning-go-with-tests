package shapes

import (
	"math"
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
		circ := Circle{1.0}
		got := Area(circle)
		want := math.Pi * math.Pow(circ.Radius, 2)

		if got != want {
			t.Errorf("got %.1f want %1.f given %v", got, want, rec)
		}
	})
}
