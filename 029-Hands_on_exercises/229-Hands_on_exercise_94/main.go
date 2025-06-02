package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"runtime"
)

func main() {
	n1 := make (chan int)
	n2 := make (chan int)

	go populate(n1)

	go fanOutIn(n1, n2)

	for v := range n2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(n chan int) {
	for i := 0; i < 100; i++ {
		n <- i
	}
	close(n)
}

func fanOutIn(n1, n2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range n1 {
				func(v2 int) {
					n2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	fmt.Println("Go Routines:",runtime.NumGoroutine())
	wg.Wait()
	close(n2)
}

func timeConsumingWork(t int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return t + rand.Intn(1000)
}