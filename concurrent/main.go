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

	c := make(chan int)
	go prod(1, 2, c)
	p1 := <-c
	go prod(3, 4, c)
	p2 := <-c
	fmt.Println("4! = ", p1*p2)

}
