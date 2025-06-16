package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}