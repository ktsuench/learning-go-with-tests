package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 5.0}
	got := Perimeter(rec)
	want := 30.0

	if got != want {
		t.Errorf("got %.1f want %.1f given %.1f and %v", got, want, recc)
	}
}

func TestArea(t *testing.T) {
	rec := Rectangle{10.0, 5.0}
	got := Area(rec)
	want := 50.0

	if got != want {
		t.Errorf("got %.1f want %1.f given %v", got, want, rec)
	}
}
