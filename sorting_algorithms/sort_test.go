package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{22, 44, 66, 11, 33, 55}, []int{11, 22, 33, 44, 55, 66}},
	}

	for _, table := range tables {
		result := bubbleSort(table.input)

		if !testEq(result, table.expected) {
			t.Errorf("The array order is not correct, got: %d, want: %d.", result, table.expected)
			return
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{22, 44, 66, 11, 33, 55}, []int{11, 22, 33, 44, 55, 66}},
	}

	for _, table := range tables {
		result := selectionSort(table.input)

		if !testEq(result, table.expected) {
			t.Errorf("The array order is not correct, got: %d, want: %d.", result, table.expected)
			return
		}
	}
}

func testEq(a, b []int) bool {

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
