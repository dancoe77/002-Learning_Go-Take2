package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}

	fmt.Println("######################################")

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		//log.Fatalln(err)
		//panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

//Println formats using the default formats for its operands and writes to standard output