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
}

func TestSumAll(t *testing.T) {
	listA := []int{1, 2}
	listB := []int{0, 9}

	got := SumAll(listA, listB)
	want := []int{3, 9}

	if got != want {
		t.Errorf("got %v want %v, given %v and %v", got, want, listA, listB)
	}
}
