package main

import "fmt"

func main() {
	anonymous()

	x := func() {
		fmt.Println("This is the second anonymous func that has run")
	}

	y := func (i int) {
		fmt.Println("How many times did the anonymous func run?", i)
	}

	x()
	y(3)
}

func anonymous() {
	fmt.Println("Anonymous ran")
}