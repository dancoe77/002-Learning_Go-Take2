package main

import "fmt"

func main() {
	var a int = 1
	fmt.Println(&a)
	/*
	./main.go:7:13: inlining call to fmt.Println
	./main.go:13:13: inlining call to fmt.Println
	./main.go:6:6: moved to heap: a
	./main.go:7:13: ... argument does not escape
	./main.go:13:13: ... argument does not escape
	./main.go:13:14: *(&a) escapes to heap
	*/
	fmt.Println(*&a)
}

/*
escapes to the heap
variable shared between main() and Println()

moved to heap
variable moved to the heap
*/