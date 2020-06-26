package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {

	unsorted := []int{22, 44, 11, 66, 33, 55}
	run(unsorted, bubbleSort)
	run(unsorted, selectionSort)
	run(unsorted, insertionSort)
	run(unsorted, shellSort)
}

func run(unsorted []int, algorithm func([]int) []int) {
	unsortedCopy := make([]int, len(unsorted))
	copy(unsortedCopy, unsorted)
	result := algorithm(unsortedCopy)
	fmt.Println(GetFunctionName(algorithm) + " result:")
	fmt.Println(result)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
