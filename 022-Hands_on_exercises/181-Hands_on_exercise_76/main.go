package main

import "fmt"

type solar struct {
	planet string
}

func (s solar) rock() {
	fmt.Println("My name is", s.planet, "and I'm in the Solar System.")
}

func (s *solar) gas() {
	s.planet = "Saturn"
	fmt.Println("My name is", s.planet, "and I'm also in the Solar System.")
}

type moon interface {
	rock()
	gas()
}

func satellite(m moon) {
	fmt.Println("#################################################")
	m.gas()
}

func main() {
	s1 := solar{"Earth"}
	s1.rock() 	//Earth
	s1.gas()	//Saturn
	// satellite(s1) // doesn't string

	s2 := &solar{"Mars"}
	s2.rock() 		//Mars
	s2.gas()		//Saturn
	satellite(s2)  	//Saturn
}