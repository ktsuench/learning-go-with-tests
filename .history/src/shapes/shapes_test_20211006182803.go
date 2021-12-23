package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	width := 10.0
	height := 5.0
	got := Perimeter(width, height)
	want := 30.0

	if got != want {
		t.Errorf("got %f want %f given %f and %f", got, want, width, height)
	}
}
