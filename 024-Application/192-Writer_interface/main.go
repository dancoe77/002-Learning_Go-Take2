package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	fmt.Println("Hello World")
	//fmt.Println calls fmt.Fprintln from package os
	fmt.Fprintln(os.Stdout, "Hello World")
	io.WriteString(os.Stdout, "Hello World\n")
}