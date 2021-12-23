package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(lists ...[]int) (sums []int) {
	for _, numbers := range lists {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(lists ...[]int) (sums []int) {
	return
}
