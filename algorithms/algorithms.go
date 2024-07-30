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
