package main

import (
	"fmt"
	"quote"
	"word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println("########################")
	fmt.Println(word.UseCount(quote.SunAlso))
}