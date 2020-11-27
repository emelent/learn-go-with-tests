package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	setA := []int{1, 2, 3}
	setB := []int{4, 5, 6}

	got := SumAll(setA, setB)
	want := 21

	if got != want {
		t.Errorf("got %d want %d given, (%v, %v)", got, want, setA, setB)
	}

}
