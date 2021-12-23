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
	checkArea := func(t testing.TB, got float64, want float64, s interface{}) {
		t.Helper()
		if got != want {
			t.Errorf("got %.1f want %1.f given %v", got, want, s)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rec := Rectangle{10.0, 5.0}
		got := rec.Area()
		want := 50.0
		checkArea(t, got, want, rec)
	})
	t.Run("circles", func(t *testing.T) {
		circ := Circle{10}
		got := circ.Area()
		want := 314.1592653589793
		checkArea(t, got, want, circ)
	})
}
