package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWithTestify(t *testing.T) {
	tables := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{55, 44, 11, 66, 33, 22}, []int{11, 22, 33, 44, 55, 66}},
	}

	for _, table := range tables {
		result := bubbleSort(table.input)

		assert.Equal(t, table.expected, result, "they should be equal")
	}
}
