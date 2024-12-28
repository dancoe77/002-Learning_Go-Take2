package main

import (
	"fmt"
)

func main(){
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	jm := []string{"Miss", "Moneypenny", "I'm 008"}
	var s string = "##############################################"
	fmt.Println(s)
	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println(s)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
	fmt.Println(len(xxs))
	fmt.Println(cap(xxs))
	fmt.Println(s)
	for i, v := range xxs {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}
	fmt.Println(s)
}
