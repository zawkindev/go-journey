package main

import "fmt"

type Waifu interface {
	Confess() string
}

type Asuna struct {
	name string
	age  int
}

func (a *Asuna) Confess() string {
	return "スキデス"
}

func printConfession[T Waifu](object T) { // main POINT
	fmt.Println(object.Confess())
}

type Comparable interface {
	~int | ~float32 | ~float64
}

func Max[T Comparable](a, b, c T) T {
	if a > b {
		if a > c {
			return a
		}
		return c
	} else {
		if c > b {
			return c
		}
		return b
	}
}

func main() {
	girlfriend := Asuna{name: "Asuna", age: 18}
	printConfession(&girlfriend)
	fmt.Println("Max of (3,9,1) is: ", Max(3, 9, 1))
	fmt.Println("Max of (3.4,9.8,1) is: ", Max(3.4, 9.8, 1))
}
