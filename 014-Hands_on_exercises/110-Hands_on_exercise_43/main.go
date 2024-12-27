package main

import (
	"fmt"
)

func main(){
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	var s string = "#####################################"
	fmt.Println(s)
	fmt.Println(xi)
	fmt.Printf("%T\n", xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println(s)

	// blank identifier to not use a returned value
	for k, v := range xi {
		fmt.Printf("%v\t %v\t %T\n", k, v, v)
	}
	fmt.Println(s)
}