package functions

import (
	"fmt"
	"io"
	"os"
)

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

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Open(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)

	defer src.Close()
	defer dst.Close()

	return
}

func handlePanic() {
	panicHappened := recover()

	if panicHappened != nil {
		fmt.Println("Panic happened already!")
	}
}

func Division(a, b int) int {
	if b == 0 {
		panic("Zero division error!")
	}

	defer handlePanic()

	return a / b
}

func PrintSlice(sli []int) {
	for index, value := range sli {
		fmt.Printf("sli[%d] = %d", index, value)
	}
}

func ApplyIt(some_func func(int) int, x int) int {
	return some_func(x)
}
