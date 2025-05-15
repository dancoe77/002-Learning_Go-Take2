package main

import "fmt"

type person struct {
	first string
}

func visibleName(p person, s string) person {
	p.first = s
	return p
}

func invisibleName(p *person, s string) {
	p.first = s
}

func main() {
	p := person{
		first: "Jenny",
	}
	fmt.Println(p)

	invisibleName(&p, "Jen")
	fmt.Println(p)
}