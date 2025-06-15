package main

import "fmt"

func main() {
	s0	:= "You have done great work - the greatest work."
	s1	:= "You have taken steps to create a better life for yourself, and for others."
	s2	:= "As an individual improves their own life, they improve the world."
	s3	:= "The skills you are acquiring are some of the most valuable skills demanded today:"
	s4	:= "knowing how to code and knowing how to use the Go programming language."
	xs := []string{s0, s1, s2, s3, s4}
	for _, v := range xs {
		fmt.Println(v)
	}
}