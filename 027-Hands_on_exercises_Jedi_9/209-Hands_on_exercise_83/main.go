package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}
func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	//this won't work
	// saySomething(p1)
	p1.speak()

	p2 := person{
		first: "Penny",
	}

	saySomething(&p2)

	//this won't work either
	// saySomething(p2)

	p2.speak()
}