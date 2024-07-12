package main

import (
	"fmt"
)

func main() {
	const name = "Kim"
	const age = 22
	
	// const name, age = "Kim", 22
	
	// fmt.Printf("%s is %d years old.\n", name, age)

	// fmt.Printf("%s is %d years old, and the type is %T and %T \n", name, age, name, age)
	
	fmt.Printf("%v is %v years old, and the type is %T and %T \n", name, age, name, age)
}