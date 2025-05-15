package main

import "fmt"

func main() {
	var i int = 66
	fmt.Println(i)
	fmt.Println(&i)
	fmt.Printf("The type of i is %T\n", i)
	fmt.Printf("The type of &i is %T\n", &i)
	fmt.Printf("The value stored at %v is %v\n", &i, *&i)

	var y string = "Illuminati"
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Printf("The type of y is %T\n", y)
	fmt.Printf("The type of &y is %T\n", &y)
	fmt.Printf("The value stored at %v is %v\n", &y, *&y)
}