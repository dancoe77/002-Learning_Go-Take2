// Package mymath perfoms math operations on a slice of type int and returns a value of type float64
package mymath

import "sort"

// CenteredAvg computes the average of a list of numbers
// after removing the largest and the smallest values.
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi)-1)]

	n := 0
	for _, v := range xi {
		n += v
	}
	f := float64(n) / float64(len(xi))
	return f
}