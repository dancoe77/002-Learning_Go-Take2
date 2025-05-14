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

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	fmt.Println("#################################")
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.walk() 	//Henry
	d1.run()	//Rover
	// youngRun(d1) // doesn't run

	d2 := &dog{"Padget"}
	d2.walk()		//Padget
	d2.run()		//Rover
	youngRun(d2)	//Rover
}