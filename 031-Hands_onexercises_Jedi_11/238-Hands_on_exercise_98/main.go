package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat 	string
	long	string
	err		error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("square root of a negative number: %v\n", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", me}
	}
	return 42, nil
}