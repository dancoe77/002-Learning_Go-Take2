// Test files live in the same package as the code they test.
// They are named with the following convention:
// 'filename_test.go' where filename is the name
// of the file that contains the code you want to test.

package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := paradise("Bali")
	want := "My idea of paradise is Bali"
	if got != want {
		log.Fatalf("error - want %s and got %s", want, got)
	}
}

// to run your test at the terminal
// aka bash, aka shell
// go mod init <project>
// go test