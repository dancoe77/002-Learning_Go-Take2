package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long 	string
	err			error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main () {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: sqaure root of negative number: %v", f)
		return 0, &NorgateMathError{"50.229 N", "99.4656 W", nme}
	}
	// implementation
	return 42, nil
}

// see use of structs with error type in standard library:
// https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html
//
// https://pkg.go.dev/net#OpError
// https://go.dev/src/net/dial.go
// https://go.dev/src/net/net.go
//
// https://go.dev/src/encoding/json/decode.go