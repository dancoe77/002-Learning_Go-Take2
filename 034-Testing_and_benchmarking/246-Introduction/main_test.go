package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2, 3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
}

func TestMyAdd(t *testing.T) {
	if myAdd(2, 3) != 2+3 {
		t.Error("Expected", 5, "Got", myAdd(2, 3))
	}
}