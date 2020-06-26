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

func insertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		toMove := numbers[i]
		reverseIndex := i - 1

		for reverseIndex >= 0 && numbers[reverseIndex] > toMove {
			numbers[reverseIndex+1] = numbers[reverseIndex]
			reverseIndex = reverseIndex - 1
		}
		numbers[reverseIndex+1] = toMove
	}

	return numbers
}

func shellSort(numbers []int) []int {
	gap := len(numbers) / 2

	for gap > 0 {
		for i := gap; i < len(numbers); i++ {
			for j := i; j >= gap && numbers[j] < numbers[j-gap]; j = j - gap {
				swap(numbers, j, j-gap)
			}
		}
		gap /= 2
	}

	return numbers
}

func swap(array []int, this, with int) {
	replaced := array[this]
	array[this] = array[with]
	array[with] = replaced
}
