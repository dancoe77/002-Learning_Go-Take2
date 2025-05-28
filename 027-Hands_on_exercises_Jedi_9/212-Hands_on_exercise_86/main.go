package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumCPU())

	var incrementor int64

	const gs = 50
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&incrementor, 1)
			runtime.Gosched()
			fmt.Println("Incrementor:\t", atomic.LoadInt64(&incrementor))

			wg.Done()
		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines:", runtime.NumGoroutine())
	fmt.Println("count:", incrementor)
}