package main

import "fmt"

func main() {
	var i int = 66
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Printf("The type of i is %T\n", i)
	fmt.Printf("The type of &i is %T\n", &i)
}