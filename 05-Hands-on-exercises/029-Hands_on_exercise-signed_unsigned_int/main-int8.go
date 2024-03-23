package main

import "fmt"

func main(){
	var a int8 = 127
	fmt.Printf("%v is of type %T\n", a, a)
	
	var b uint8 = 255
	fmt.Printf("%v is of type %T\n", b, b)
}