package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("sum 10 numbers", func(t *testing.T) {
		numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		got := Sum(numbers)
		want := 45

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
