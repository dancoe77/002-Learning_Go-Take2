package main

import "fmt"

func main() {
	fmt.Println("2 + 3 =", mySum(2, 3))
	fmt.Println("4 + 7 =", mySum(4, 7))
	fmt.Println("5 + 9 =", mySum(5, 9))

	fmt.Println("#############################")

	fmt.Println("2 + 3 =", myAdd(2, 3))
	fmt.Println("4 + 7 =", myAdd(4, 7))
	fmt.Println("5 + 9 =", myAdd(5, 9))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func myAdd (xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}