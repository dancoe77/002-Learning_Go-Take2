package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First 	string
	Age 	int
}

// ByAge implements sort.Interface for []Person based on
// the Age field
type ByAge []Person
type ByName []Person

func (a ByAge) Len() int			{ return len(a) }
func (a ByAge) Swap(i, j int)		{ a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool	{ return a[i].Age < a[j].Age}

func (n ByName) Len() int		{ return len(n) }
func (n ByName) Swap(i, j int)		{ n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool	{ return n[i].First < n[j].First}

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println("--------------------------------------------")
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}