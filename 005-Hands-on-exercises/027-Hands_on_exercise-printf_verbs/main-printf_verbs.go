package main

import (
	"fmt"
	"math"
)

func main() {

	var a string = "knowledge"
	fmt.Printf("%v is of type %T\n", a, a)

	var b int = 69
	fmt.Printf("%v is of type %T\n", b, b)

	var c,d int = 3, 4
	var e float64 = math.Sqrt(float64(c*c + d*d))
	fmt.Printf("%v is of type %T\n", e, e)
}
//
/*
for a VARIABLE storing a VALUE of TYPE
	string
	int
	float64
use print verbs to show
	value
	type
*/