package main

import (
	"fmt"
)

func main (){
	m := map[string]int{
		"James":		42,
		"Moneypenny":	32,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}