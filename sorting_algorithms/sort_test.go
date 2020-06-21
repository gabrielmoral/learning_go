package main

import (
	"testing"
)

func TestSortAlgorithms(t *testing.T) {
	algorithms := []func([]int) []int{bubbleSort, selectionSort, insertionSort, shellSort}

	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{22, 44, 66, 11, 33, 55}, []int{11, 22, 33, 44, 55, 66}},
		{[]int{22, 44, 66, 9, 11, 33, 111, 110, 55}, []int{9, 11, 22, 33, 44, 55, 66, 110, 111}},
	}

	for _, table := range tables {
		for _, algorithm := range algorithms {
			result := algorithm(table.input)

			if !IsEq(result, table.expected) {
				t.Errorf("The array order is not correct, got: %d, want: %d. with %s", result, table.expected, GetFunctionName(algorithm))
				return
			}
		}
	}
}

func IsEq(a, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
