package main

import (
	"fmt"
	"math/rand"
)

func main(){
	
	x := rand.Intn(10)
	
	for x < 10 {
		fmt.Printf("x is %v\n", x)
		if x < 2 {
			break
		}
		x++
	}
}