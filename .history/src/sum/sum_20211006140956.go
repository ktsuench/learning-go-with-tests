package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(lists ...[]int) (results []int) {
	return
}
