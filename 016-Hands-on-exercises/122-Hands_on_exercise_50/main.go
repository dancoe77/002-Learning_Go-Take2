package main

import "fmt"

func main() {
	am := make(map[string][]string)
	am["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	am["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	am["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	am["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	fmt.Println(len(am))

	for k, v := range am {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}