package main

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(lists ...[]int) (sums []int) {
	for _, nums := range lists {
		sums = append(sums, Sum(nums))
	}
	return
}

func SumAllTails(lists ...[]int) (sums []int) {
	return
}
