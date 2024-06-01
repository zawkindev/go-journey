package main

import (
	"fmt"
	f "go-journey/functions"
)

func main() {

	f.Hello()
	
	x, y := f.Calculate(3, 4)
	fmt.Printf("area: %d, perimeter: %d\n", x, y)

	fmt.Printf("x: %d, y: %d", x, y)
	f.Swap(&x, &y)
	fmt.Printf("\n---Swapped---\n")
	fmt.Printf("x: %d, y: %d", x, y)

	f.Say(func(txt string) string {
		return txt
	}("salom"),
	)

	firstSum := f.PartialSum(5)
	fmt.Printf("\nPARTIAL SUM: %d", firstSum(5))

	fmt.Printf("\nTRIPLEsum func: %d\n", f.TripleSum(3)(4)(7))

	f.VariadicFunc(1, 2, 3, 4, 5, 6)
	f.WithDifferentTypes("salom", 9879087, 4.323)

}
