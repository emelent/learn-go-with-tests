package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(sets ...[]int) int {
	sum := 0
	for _, set := range sets {
		sum += Sum(set)
	}
	return sum
}
