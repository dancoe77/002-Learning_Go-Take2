package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go routines\t", runtime.NumGoroutine())

	wg.Add(2)
	go nebula()
	go blackhole()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Go routines\t", runtime.NumGoroutine())
	wg.Wait()
}

func nebula() {
	fmt.Println("Nebula ran")
	wg.Done()
}

func blackhole() {
	fmt.Println("Blackhole ran")
	wg.Done()
}

