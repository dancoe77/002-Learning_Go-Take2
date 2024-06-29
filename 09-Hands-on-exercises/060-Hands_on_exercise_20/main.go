package main

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
)

func main(){
	p1 := puppy.Bark()
	p2 := puppy.Barks()
	fmt.Println(p1)
	fmt.Println(p2)
}