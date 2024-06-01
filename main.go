package main

import (
	"fmt"
	f "go-journey/functions"

	a "go-journey/arrays"
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

	fmt.Println("\n##################Arrays###################")
	a.TenZero()

	var array [2]string = [2]string{"hi", "bye"}
	fmt.Println(array)

	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	var sl = []int{3, 1, 1, 4}
	fmt.Println(sl[0:2])

	fmt.Println("len: ", len(sl))

	fmt.Println("cap: ", cap(sl))

	fmt.Println("sl from arr: ", arr[0:2])

	var msl = make([]int, 2, 8)

	fmt.Println("msl len: ", len(msl))

	fmt.Println("msl cap: ", cap(msl))

	msl = append(msl, 20, 21)

	fmt.Println("after append: ", msl)

	new_msl := []int{12, 123123, 121233}

	msl = append(msl, new_msl...)

	fmt.Println("appending two slices: ", msl)

	neededNums := msl[:len(msl)-3]
	dest := make([]int, len(neededNums))

	copy(neededNums,dest)

	fmt.Println("copied nums: ", dest)
}
