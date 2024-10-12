package main

import (
	"fmt"
)

func main(){
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream",
		"Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)
	fmt.Println(len(xs))

	// blank identifier to not use a returned value
	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}
	var s string = "#####################################"
	fmt.Println(s)

	fmt.Println(xs[0])
	fmt.Println(xs[10])
	fmt.Println(xs[20])
	fmt.Println(xs[30])

	// doesn't work
	//fmt.Println(xs[31])

	fmt.Println(len(xs))

	fmt.Println(s)

	for i := 0; i<len(xs); i++ {
		fmt.Println(xs[i])
	}
}