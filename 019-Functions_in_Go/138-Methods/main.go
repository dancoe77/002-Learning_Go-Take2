package main

import "fmt"

type person struct {
	first string
	last string
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
	}

	p2 := person{
		first: "Jenny",
		last: "Moneypenny",
	}

	p1.speak()
	p2.speak()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }