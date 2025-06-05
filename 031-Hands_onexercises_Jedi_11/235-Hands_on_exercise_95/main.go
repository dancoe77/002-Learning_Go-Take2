package main

import (
	"fmt"
	"encoding/json"
	"os"
	"log"
)

type person struct {
	First 	string
	Last	string
	Sayings	[]string
}

func main() {
	p1 := person {
		First:		"James",
		Last:		"Bond",
		Sayings:	[]string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	f, err := os.Create("error.log")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println("err happened", err)
	}
	fmt.Println(string(bs))
	log.Println(string(bs))
}