package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(7, 5)
	want := 2
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(7, 5, add)
	want := 12
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d", got, want)
	}
}