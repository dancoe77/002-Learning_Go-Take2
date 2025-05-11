package main

import "fmt"

func main(){
	x := anonymous()
	fmt.Println(x)

	y := illuminati()
	fmt.Println(y())

	fmt.Printf("%T\n", anonymous)
	fmt.Printf("%T\n", illuminati)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

func anonymous() int {
	return 42
}

func illuminati() func() string {
	return func() string {
		return "Walk the Path of Illumination"
	}
}