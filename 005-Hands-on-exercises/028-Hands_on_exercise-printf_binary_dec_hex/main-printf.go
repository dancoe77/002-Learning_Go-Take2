package main

import "fmt"

func main() {
	a, b, c := 747, 911, 90210
	fmt.Println("Value \t Binary \t\t Hex")
	fmt.Printf("%v \t %b \t\t %#x \n",a,a,a)
	fmt.Printf("%v \t %b \t\t %#x \n",b,b,b)
	fmt.Printf("%v \t %b \t %#x \n",c,c,c)
}

//
/*
print verbs to show the following numbers
	747
	911
	90210
in
	decimal
	bimary
	hex

*/