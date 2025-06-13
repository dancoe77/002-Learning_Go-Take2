package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
//Only necessary prior to go 1.5 with 1.5 or past multiple cores are automatically used
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
*/

func main() {
	// No concurrency currently
	// foo()
	// bar()

	// Concurrency
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}