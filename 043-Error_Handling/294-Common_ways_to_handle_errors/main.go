package main

import (
	//"fmt"
	//"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//				fmt.Println("err happened", err)
		//				log.Println("err happened", err)
		//				log.Fatalln(err)
		panic(err)
	}
}

// Package log implements a simple logging package ...
// writes to standard err and prints the date and time of each logged message ...
// the Fatal functions call os.Exit(1) after writing the log message ...
// the Panic functions call panic after writing the log message.

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).