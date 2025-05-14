package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

func main() {
	d1 := dog{"Henry"}
	d1.walk() 	//Henry
	d1.run()	//Rover

	d2 := &dog{"Padget"}
	d2.walk()	//Padget
	d2.run()	//Rover
}