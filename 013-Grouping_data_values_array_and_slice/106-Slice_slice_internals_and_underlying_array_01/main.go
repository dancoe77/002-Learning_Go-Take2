package main

import (
	"fmt"
)

func main(){
	a := []int{0, 1, 2, 3, 4, 5}
	b := a
	var s string = "##############################"
	fmt.Println(s)
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println(s)

	a[0] = 7
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println(s)
}