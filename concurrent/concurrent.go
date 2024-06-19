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

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printToN("goroutine: ", 1000, &wg)
	wg.Wait()

	for i := 0; i < 5; i++ {
		fmt.Println("main: ", i)
	}
}
