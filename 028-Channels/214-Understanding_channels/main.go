package main

import "fmt"

func main(){
	c := make(chan int)

	go func(){
		c <- 42
	}()
	fmt.Println(<- c)
}

/* This will not run
c := make(chan int)
c <- 42
fmt.Println(<- c)
because channels block and
there is nothing to pull
the value off immediately
*/

/* This will run successfully
because of the buffer
c := make(chan int, 1)
c <- 42
fmt.Println(<- c)