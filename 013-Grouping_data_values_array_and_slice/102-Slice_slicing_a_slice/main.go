package main

import (
	"fmt"
)

func main(){
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s string = "##############################"
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(s)

	// [inclusive:exclusive]
	fmt.Printf("xi - %#v\n", xi[0:4])
	fmt.Println(s)

	// [:exclusive]
	fmt.Printf("xi - %#v\n", xi[:7])
	fmt.Println(s)

	// [inclusive:]
	fmt.Printf("xi - %#v\n", xi[4:])
	fmt.Println(s)

	// [:]
	fmt.Printf("xi - %#v\n", xi[:])
	fmt.Println(s)
}