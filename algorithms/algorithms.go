package algorithms

import "fmt"

func BinarySearch(list []int, n int) int {
	low, high := 0, len(list)-1

	for low <= high {
		median := (low + high) / 2
		current := list[median]

		switch {
		case current == n:
			return median
		case current < n:
			low = median + 1
		case current > n:
			high = median - 1
		}
	}

	return -1
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	n := BinarySearch(list, 3)
	fmt.Println("index: ", n)
}
