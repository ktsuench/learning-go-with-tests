package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(lists ...[]int) (sums []int) {
	for i := 0; i < 2; i++ {
		sums = append(sums, Sum(lists))
	}
	return
}
