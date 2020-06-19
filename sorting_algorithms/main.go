package main

import (
	"fmt"
)

func main() {

	numbers := []int{22, 44, 11, 66, 33, 55}
	fmt.Println("original              ", numbers)

	bubbleSortResult := bubbleSort(numbers)
	selectionSortResult := selectionSort(numbers)

	fmt.Println("bubble sort result    ", bubbleSortResult)
	fmt.Println("selection sort result ", selectionSortResult)
}
