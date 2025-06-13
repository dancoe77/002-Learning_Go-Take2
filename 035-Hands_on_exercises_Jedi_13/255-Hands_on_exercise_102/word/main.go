// Package word performs operations on a string input and returns an int
package word

import (
	"strings"
	"maps"
	"slices"
)
// UseCount takes in a value of type string, creates a map of string and int, imports the values into a slice of type int and then adds those values together to return a type of value int
func UseCount(s string) int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	sum := 0
	for _, v := range xs {
		m[v]++
	}
	xm := slices.Collect(maps.Values(m))
	numbers := []int(xm)
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Count takes in a value of type string and returns a value of type int
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}