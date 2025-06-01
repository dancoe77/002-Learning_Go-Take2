package main

import "fmt"

func main(){
	e := make(chan int)
	o := make(chan int)
	q := make(chan int)

	//send
	go gen(e, o, q)

	//receive
	receive(e, o, q)

	fmt.Println("about to exit")
}

//receive
func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <- even:
			fmt.Println("from the even channel:", v)
		case v := <- odd:
			fmt.Println("from the odd channel", v)
		case v := <- quit:
			fmt.Println("from the quit channel:", v)
			//close(quit)
			return
		}
	}
}

//send
func gen(even, odd, quit chan<- int){
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	quit <- 0
}
