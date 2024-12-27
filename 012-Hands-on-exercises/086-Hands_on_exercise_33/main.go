package main

import (
	"fmt"
)

func main(){
	for i := 0; i <= 11; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("counting odd numbers: ", i)
	}
}