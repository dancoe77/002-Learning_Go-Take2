package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go incrementor(c)
	for n := range c {
		fmt.Println(n)
	}
}

func incrementor(c chan int) {
	for i := 0; i < 20; i++ {
			c <- i
	}
	close(c)
}