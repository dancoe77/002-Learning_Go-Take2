package main

import (
	"fmt"
	"math/rand"
)

func main(){
	rand.Intn(x); x <= 250
	fmt.printf("x is %v \n", x)

	if x <= 100 {
		fmt.Println("x is between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x is between 101 and 200")
	} else {
		fmt.Println("x is between 201 and 250")
	}
	/*Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n) from the default Source.
	It panics if n <= 0.  [func Intn](https://pkg.go.dev/math/rand#Intn)*/
}