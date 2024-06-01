package main

import "fmt"

func calculate(a, b int) (area, perimeter int) {
	area = a * b
	perimeter = (a + b) * 2
	return
}

func swap(x, y *int) {
	var temp int = *x
	*x = *y
	*y = temp
}

var say = func(txt string) {
	fmt.Printf("%s", txt)
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func tripleSum(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func main() {
	x, y := calculate(3, 4)
	fmt.Printf("area: %d, perimeter: %d\n", x, y)

	fmt.Printf("x: %d, y: %d", x, y)
	swap(&x, &y)
	fmt.Printf("\n---Swapped---\n")
	fmt.Printf("x: %d, y: %d", x, y)

	say(func(txt string) string {
		return txt
	}("salom"),
	)

	firstSum := partialSum(5)
	fmt.Printf("\nPARTIAL SUM: %d", firstSum(5))

	fmt.Printf("\nTRIPLEsum func: %d", tripleSum(3)(4)(7))
}
