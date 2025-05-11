package main

import "fmt"

func main() {
	anonymous()

	func() {
		fmt.Println("This is the second anonymous func that has run")
	}()

	func (i int) {
		fmt.Println("How many times did the anonymous func run?", i)
	}(3)
}

func anonymous() {
	fmt.Println("Anonymous ran")
}