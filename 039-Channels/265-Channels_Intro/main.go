package main

import (
	"fmt"
	"time"
)

func main(){
	proverbs()
	fmt.Println("##############################################################################")
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}

func proverbs(){
	s0 	:= "Go Proverbs"
	s1 	:= "#################################################################################"
	s2 	:= "Do not communicate by sharing memory; instead, share memory by communicating."
	s3 	:= "Concurrency is not parallelism."
	s4 	:= "Channels orchestrate; mutexes serialize."
	s5 	:= "The bigger the interface, the weaker the abstraction."
	s6 	:= "Make the zero value useful."
	s7 	:= "interface{} says nothing."
	s8 	:= "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite."
	s9 	:= "A little copying is better than a little dependency."
	s10 := "Syscall must always be guarded with build tags."
	s11	:= "Cgo must always be guarded with build tags."
	s12	:= "Cgo is not Go."
	s13	:= "With the unsafe package there are no guarantees."
	s14	:= "Clear is better than clever."
	s15	:= "Reflection is never clear."
	s16	:= "Errors are values."
	s17	:= "Don't just check errors, handle them gracefully."
	s18	:= "Design the architecture, name the components, document the details."
	s19	:= "Documentation is for users."
	s20	:= "Don't panic."
	xs := []string{s0, s1, s2, s3, s4, s5, s6, s7, s8, s9, s10, s11, s12, s13, s14, s15, s16, s17, s18, s19, s20}
	for _, v := range xs {
		fmt.Println(v)
	}
}