package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("repeat 'b' 5 times", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
