package algorithms

import "fmt"

func BinarySearch(list []int, n int) (int, error) {
	low, high := 0, len(list)-1

	if list[low] > list[high] {
		return -1, fmt.Errorf("unsorted array")
	}

	for low <= high {
		median := (low + high) / 2
		current := list[median]

		switch {
		case current == n:
			return median, nil
		case current < n:
			low = median + 1
		case current > n:
			high = median - 1
		}
	}

	return -1, fmt.Errorf("algoritm failed")
}

func SelectionSort(list []int) []int {
	newList := make([]int, len(list))
	n := len(list)

	for i := 0; i < n; i++ {
		maxIndex := 0
		for j := 1; j < len(list); j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
			}
		}
		newList[i] = list[maxIndex]
		list = append(list[:maxIndex], list[maxIndex+1:]...)
	}
	return newList
}

func Factorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * Factorial(n-1)
	}
}

func EKUB(a, b int) int { // The Euclidean algorithm
	if b == 0 {
		return a
	} else {
		return EKUB(b, a%b)
	}
}
