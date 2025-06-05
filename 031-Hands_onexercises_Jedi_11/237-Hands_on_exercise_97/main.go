package main

import (
	"fmt"
	"log"
)

/*type error interface {
	Error() string
}*/

type customErr struct {
	level	string
	err 	error
}

func (c customErr) Error() string {
	return fmt.Sprintf("[%v] %v\n", c.level, c.err)
}

func main() {
	_, err := foo("Hello Dan")
	if err != nil {
		log.Println(err)
	}
}

func foo(s string) (string, error) {
	ce := fmt.Errorf("%v the application has stopped", s)
	return "nil", customErr{"info", ce}
}