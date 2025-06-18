package main

import "fmt"

func main() {
	f := factorial(4)
	for n := range f {
		fmt.Println("Total:", n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}