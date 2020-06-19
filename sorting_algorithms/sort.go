package main

func bubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				swap(numbers, i, j)
			}
		}
	}

	return numbers
}

func selectionSort(numbers []int) []int {
	for wall := len(numbers) - 1; wall > 0; wall-- {
		largestAt := 0
		for j := 0; j <= wall; j++ {

			if numbers[j] > numbers[largestAt] {
				largestAt = j
			}
		}
		swap(numbers, wall, largestAt)
	}

	return numbers
}

func swap(array []int, this, with int) {
	replaced := array[this]
	array[this] = array[with]
	array[with] = replaced
}
