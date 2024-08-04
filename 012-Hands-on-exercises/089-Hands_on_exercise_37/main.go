package main

import (
	"fmt"
)

func main (){
	var bond = map[string]int{
		"James":		42,
		"Moneypenny":	32,
		//"Q":			52,
	}
	for k, v := range bond {
		fmt.Println("ranging over a map", k, v)
	}
	m1 := bond["James"]
	fmt.Println("The age of Bond is", m1)
	
	if v, ok := bond["James"]; ok {
		fmt.Println("There is a Bond lookup entry, and Bond's age is", v)
	}

	m2 := bond["Q"]
	fmt.Println("The age of Q is", m2)
	
	if v, ok := bond["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int", v)
	}
}