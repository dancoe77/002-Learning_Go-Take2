package main

import (
	"fmt"
)

func main(){
	si := []string{"a", "b", "c"}
	var s string = "##############################"
	fmt.Println(si)
	fmt.Println(s)

	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println(s)
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println(s)


}
/*
xi := make([]int, 0, 10)
fmt.Println(xi)
fmt.Println(len(xi))
fmt.Println(cap(xi))
xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
var s string = "##############################"
fmt.Printf("xi - %#v\n", xi)
fmt.Println(s)
*/