package main

import "fmt"

func main() {
	n := make(chan int)

	go send(n)

	receive(n)

	fmt.Println("about to exit")
}

//send channel
func send(n chan<- int){
	for i := 0; i < 100; i++ {
		n <- i
	}
	close(n)
}

//receive
func receive(n <-chan int) {
	for v := range n {
		fmt.Println(v)
	}
}