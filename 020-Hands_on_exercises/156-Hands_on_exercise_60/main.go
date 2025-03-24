package main

import "fmt"

func main() {
	defer foo()
	bar()

	defer planet()
	cosmos()
}

func foo() {
	fmt.Println("foo ran 1st")
}

func bar() {
	fmt.Println("bar ran 2nd")
}

func planet() {
	fmt.Println("Hello World ran 3rd")
}

func cosmos() {
	fmt.Println("Hello Galaxy ran 4th")
}