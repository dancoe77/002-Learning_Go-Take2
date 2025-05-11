package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("The type of x is %T\n", x)
	fmt.Printf("The type of &x is %T\n", &x)
}