package main

import "fmt"

type person struct {
	title string
	first string
	middle string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("I am", p.title, p.first, p.middle, p.last, "and my age is", p.age)
}

func main() {
	p1 := person{
		title: "Captain",
		first: "James",
		middle: "Tiberius",
		last: "Kirk",
		age: 42,
	}

	p2 := person{
		title: "Captain",
		first: "Jean",
		middle: "Luc",
		last: "Picard",
		age: 43,
	}

	p3 := person{
		title: "First Officer",
		first: "William",
		middle: "Thomas",
		last: "Riker",
		age: 44,
	}

	p1.speak()
	p2.speak()
	p3.speak()
}