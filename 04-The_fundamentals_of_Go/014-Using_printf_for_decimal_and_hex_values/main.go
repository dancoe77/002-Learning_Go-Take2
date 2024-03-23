package main

import "fmt"

func main(){
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)
	
	//print these values as both binary & hex
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Println("Value \t Binary  Hex")
	fmt.Printf("%v \t %b \t %#x \t \n",a,a,a)
	fmt.Printf("%v \t %b \t %#x \t \n",b,b,b)
	fmt.Printf("%v \t %b \t %#x \t \n",c,c,c)
	fmt.Printf("%v \t %b \t %#x \t \n",d,d,d)
	fmt.Printf("%v \t %b \t %#x \t \n",e,e,e)
	fmt.Printf("%v \t %b \t %#x \t \n",f,f,f)
	
}
