package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex
var once sync.Once

func printToN(msg string, x int, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	for i := 1; i <= x; i++ {
		fmt.Println(msg, i)
	}
	mutex.Unlock()
}

func prod(a, b int, c chan int) {
	c <- a * b
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("########### WaitGroup ###########")

	wg.Add(1)

	go printToN("goroutine: ", 5, &wg)

	wg.Wait()

	for i := 0; i < 2; i++ {
		go fmt.Println("main: ", i)

		once.Do(func() {
			fmt.Println("init")
		})
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

	fmt.Println("########### Context ###########")

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(ctx context.Context) {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Completed Successfully!!!")
		case <-ctx.Done():
			fmt.Println("Canceled or timed out: ", ctx.Err())
		}
	}(ctx)

	time.Sleep(2*time.Second)
}
