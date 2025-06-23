package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//				fmt.Println("err happened", err)
		log.Println("err happened", err)
		//				log.Fatalln(err)
		//				panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.

// Package log implements a simple logging package ...
// writes to standard err and prints the date and time of each logged message ...
// the Fatal functions call os.Exit(1) after writing the log message ...
// the Panic functions call panic after writing the log message.

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println