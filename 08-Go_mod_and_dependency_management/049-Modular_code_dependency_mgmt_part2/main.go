package main

import(
	"fmt"
	"github.com/dancoe77/008-puppy"
	"github.com/dancoe77/009-dog"
)
func main(){
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
	
	s3 := dog.BigBark()
	s4 := dog.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)
}