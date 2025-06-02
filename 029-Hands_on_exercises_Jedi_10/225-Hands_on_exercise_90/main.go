package main

import "fmt"

func main(){
	c := make(chan int)
	go gen(c)
	receive(c)
	fmt.Println("about to exit")
}
//send
func gen(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
//receive
func receive(c <-chan int){
	for v := range c {
		fmt.Println(v)
	}
}