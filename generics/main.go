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

func printConfession[T Waifu](object T) {
	fmt.Println(object.Confess())
}

func main() {
	girlfriend := Asuna{name: "Asuna", age: 18}
	printConfession(&girlfriend)
}
