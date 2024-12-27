package main

import (
	"fmt"
)

func main(){
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Merlot", "Strawberry"}
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

}
