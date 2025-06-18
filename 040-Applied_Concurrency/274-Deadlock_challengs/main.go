package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

	fmt.Println("##########################")

	h := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			h <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-h)
		}
	}()
	time.Sleep(time.Second)
}