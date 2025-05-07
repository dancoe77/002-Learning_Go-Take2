package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(6, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", total, 10)
	}
}