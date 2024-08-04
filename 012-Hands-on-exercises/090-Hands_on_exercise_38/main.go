package main

import (
	"fmt"
	"math/rand"
)

func main(){
	for i := 0; i < 100; i++ {
		fmt.Printf("%v: ", i)
		x := rand.Intn(5)
		fmt.Printf("x is %v \n", x)

		switch {
		case x <= 2:
			fmt.Println("x is less than or equal to 2")
			case x == 3:
				fmt.Println("x is 3")
		case x > 3:
			fmt.Println("x is greater than 3")
		}
	}
}	