package main

import (
	"fmt"
)

func main(){
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s string = "##############################"
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(s)

	xi = append (xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println(s)
}