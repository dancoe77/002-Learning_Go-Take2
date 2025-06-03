package main

import (
	"fmt"
	"os"
	"strings"
	"io"
	"io/ioutil"
)

func main() {
	fmt.Println("#########################################")
	n, err := fmt.Println("Hello World")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	fmt.Println("#########################################")

	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)

	fmt.Println("#########################################")

	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("James Bond")

	io.Copy(f, r)

	//#########################################

	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}