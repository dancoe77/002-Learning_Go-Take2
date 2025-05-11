package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2)
	fmt.Println(x()) //1
	fmt.Println(x()) //2
	fmt.Println(x()) //3
	fmt.Println(x()) //4
	fmt.Println(x()) //5
	fmt.Println(x()) //6
	fmt.Println(x()) //7
	fmt.Println(x()) //8
}

func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}