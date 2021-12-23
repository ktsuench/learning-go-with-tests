package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	width := 10.0
	height := 5.0
	got := Perimeter(width, height)
	want := 30.0

	if got != want {
		t.Errorf("got %.1f want %.1f given %.1f and %.1f", got, want, width, height)
	}
}

func TestArea(t *testing.T) {
	width := 10.0
	height := 5.0
	got := Area(width, height)
	want := 50.0

	if got != want {
		t.Errorf("got %.1f want %1.f given %.1f and %.1f", got, want, width, height)
	}
}
