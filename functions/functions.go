package functions

import "fmt"

func Hello() {
	fmt.Println("hello world")
}

func Calculate(a, b int) (area, perimeter int) {
	area = a * b
	perimeter = (a + b) * 2
	return
}

func Swap(x, y *int) {
	var temp int = *x
	*x = *y
	*y = temp
}

var Say = func(txt string) {
	fmt.Printf("%s", txt)
}

func PartialSum(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func TripleSum(a int) func(int) func(int) int {
	return func(b int) func(int) int {
		return func(c int) int {
			return a + b + c
		}
	}
}

func VariadicFunc(s ...int) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d\n", s[i])
	}
}

func WithDifferentTypes(list ...interface{}) {
	for _, el := range list {
		fmt.Println(el)
	}
}
