package main

func Sum(numbers [5]int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
