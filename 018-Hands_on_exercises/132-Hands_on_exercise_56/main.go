package main

import "fmt"

type person struct {
	first string
	friends map[string]int
	favDrinks []string
}

func main() {
	p1 := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Dan",
		friends: map[string]int{"Todd": 42, "Henry": 16, "Padget": 14},
		favDrinks: []string{"Double Chocolate Mocha", "Coca-Cola","Merlot", "Water"},
	}

	p2 := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Colista",
		friends: map[string]int{"Lucas": 28, "Steph": 37, "George": 78},
		favDrinks: []string{"Americano Hot", "Coca-Cola","Water", "Moscato"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)

	fmt.Println(p1.first,p1.friends,p1.favDrinks)
	fmt.Println(p2.first,p2.friends,p2.favDrinks)
}