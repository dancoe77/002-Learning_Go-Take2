package main

import (
	"fmt"
	"math/rand"
)

func main(){
	for i := 0; i <= 42; i++ {
		fmt.Printf("%v: ", i)
		x := rand.Intn(5)
		fmt.Printf("x is %v \n", x)

		switch {
		case x < 2:
			fmt.Println("x is less than 2")
			case x >= 2 && x <= 4:
				fmt.Println("x is between 2 and 4")
		}
	}
}	