package arrays

import "fmt"

func TenZero() {
	var a [10]int
	for _, item := range a {
		fmt.Println(item)
	}
}
