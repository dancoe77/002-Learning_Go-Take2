package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	n1 := make (chan int)
	n2 := make (chan int)

	go gen(n1)

	go factorial(n1, n2)

	for v := range n2 {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}

func gen(n chan int) {
	for i := 0; i < 1000; i++ {
		n <- i
	}
	close(n)
}

func factorial(n1, n2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 100
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range n1 {
				func(v2 int) {
					n2 <- fact(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	wg.Wait()
	close(n2)
}

func fact(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH THE SOLUTION BEFORE DOIN CHALLENGE #2
-- While running the factorial computations, try to find how much of my resources are being used
-- Post the percentage of my resources being used in this comment space
*/