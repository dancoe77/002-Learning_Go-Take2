package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!!")
	
	
	//raw string literal
	fmt.Println(`
UTF-8 saves space,
In UTF-8, common characters like "C" take 8 bits,
while rare characters like "heart emoji" take 32 bits.
Other characters take 16 or 24 bits.
A blog post like this one takes about
four times less space in UTF-8
than it would in UTF-32.
So it loads four times faster.
`)

//https://go.dev/ref/spec#String_literals
}