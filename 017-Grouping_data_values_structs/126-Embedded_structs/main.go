package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
	first string
}

func main() {
	sa1 := secretAgent {
		person: person{
			first:	"James",
			last:	"Bond",
			age:	32,
			},
			ltk: true,
			first: "Ethan Hawke",
	}

	sa2 := secretAgent {
		person: person{
			first:	"Jenny",
			last:	"Moneypenny",
			age:	27,
			},
			ltk: false,
			first: "Jessica Henwick",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)
	fmt.Printf("%T \t %#v\n", sa2, sa2)

	fmt.Println(sa1.first,sa1.person.first, sa1.person.last, sa1.age, sa1.ltk)
	fmt.Println(sa2.first, sa2.person.first, sa2.person.last, sa2.age, sa2.ltk)
}