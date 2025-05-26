package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go routines:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))

			wg.Done()
		}()
		fmt.Println("Go routines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}