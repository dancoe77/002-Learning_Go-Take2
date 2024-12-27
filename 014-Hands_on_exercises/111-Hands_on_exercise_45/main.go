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

	xj = append(xi, 52)
	fmt.Println(xj)
	fmt.Println(s)

	xk = append(xj, 53, 54, 55)
	fmt.Println(xk)
	fmt.Println(s)
}
