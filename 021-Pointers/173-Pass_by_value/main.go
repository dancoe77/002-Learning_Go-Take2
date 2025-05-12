package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
	ii[1] = 100
	ii[2] = 101
	ii[3] = 102
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1,2,3,4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])
}