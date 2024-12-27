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

	xi = append(xi, 52)
	fmt.Println(xi)
	fmt.Println(s)

	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)
	fmt.Println(s)

	yi := []int{56, 57, 58, 59, 60}
	zi := append(xi, yi...)
	fmt.Println(zi)
	fmt.Println(s)
}
