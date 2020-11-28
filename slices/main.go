package main

// Sum returns the sum of a slice of integers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll2 returns the sum of a set of slices
func SumAll2(sets ...[]int) int {
	sum := 0
	for _, set := range sets {
		sum += Sum(set)
	}
	return sum
}

// SumAll returns a slice containing the individual sums of a set of slices
func SumAll(sets ...[]int) (sums []int) {
	for _, set := range sets {
		sums = append(sums, Sum(set))
	}
	return
}

func SumAllTails(sets ...[]int) (sums []int) {
	for _, set := range sets {
		if len(set) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(set[1:]))
		}
	}
	return
}
