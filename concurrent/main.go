package main

import (
	"fmt"
	"sync"
)

func printToN(msg string, x int, wg *sync.WaitGroup) {
	for i := 1; i <= x; i++ {
		fmt.Println(msg, i)
	}
	wg.Done()
}

func prod(a, b int, c chan int) {
	c <- a * b
}

func main() {
	fmt.Println("########### WaitGroup ###########")
	var wg sync.WaitGroup

	wg.Add(1)
	go printToN("goroutine: ", 3, &wg)
	wg.Wait()

	for i := 0; i < 2; i++ {
		fmt.Println("main: ", i)
	}

	fmt.Println("########### Channels ###########")

	var p1, p2 int

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	go prod(1, 2, c1)
	go prod(3, 4, c2)

	select {
	case p1 = <-c1:
		fmt.Println("p1 recieved first")
	case p2 = <-c2:
		fmt.Println("p2 recieved first")
	}

	fmt.Println("4! = ", p1*p2)

}
