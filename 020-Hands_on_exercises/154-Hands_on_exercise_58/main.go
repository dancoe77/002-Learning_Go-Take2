package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
		return 13, "Hello World"
}