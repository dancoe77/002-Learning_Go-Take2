package main

import (
	"fmt"
)

func main(){
	xi := []int{42, 43, 44}
	fmt.Println(xi)

	var s string = "##############################"
	fmt.Println(s)

	//variadic parameter
	xi = append(xi, 45, 46, 47, 99, 000)
	fmt.Println(xi)
	fmt.Println(s)
}