package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first, sa.last, sa.ltk)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Jenny",
			last: "Moneypenny",
		},
		ltk: false,
	}

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

	saySomething(sa1)
	saySomething(sa2)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }