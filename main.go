package main

import (
	"fmt"
	f "go-journey/functions"

	a "go-journey/arrays"
)

func main() {

	//Functions and Packages

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

	// Arrays

	fmt.Println("\n################## Arrays ###################")
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

	// Slices

	fmt.Println("\n################## Slices ###################")

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

	copy(neededNums, dest)
	fmt.Println("copied nums: ", dest)

	// Maps

	fmt.Println("\n################## Maps ###################")

	var strMap = map[string]string{"new": "yangi"}
	fmt.Println("strMap: ", strMap)

	var intStrMap = map[int]string{1: "bir", 2: "ikki", 3: "uch"}
	fmt.Println("intStrMap: ", intStrMap)

	var emptyMap = make(map[string]string)
	fmt.Println("emptyMap: ", emptyMap)

	emptyMap["salom"] = "hi"
	fmt.Println("after adding salom: ", emptyMap)

	delete(emptyMap, "salom")
	fmt.Println("again empty: ", emptyMap)

	val, ok := intStrMap[1]
	fmt.Printf("%s in intStrMap: %t\n", val, ok)

	fmt.Println("iterate intStrMap:")
	for k, v := range intStrMap {
		fmt.Printf("%d: %s\n", k, v)
	}

	fmt.Println("\n################## Structs ###################")

	type waifu struct {
		name string
		age  int
	}

	var megumi waifu = waifu{name: "Megumi", age: 18}
	megumi.name = "Asuna"
	fmt.Println(megumi)
}
