package main

import (
	"fmt"
	"log"
	//"errors"
)

//var ErrNorgateMath = errors.New("norgate math: square root of negative number")

type norgateMathError struct {
	lat 	string
	long	string
	err		error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	//fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}