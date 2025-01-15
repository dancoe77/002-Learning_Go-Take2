package main

import "fmt"

type person struct {
	first string
	last string
	icecream []string
}

func main(){
p1 := person{
	first: 		"Dan",
	last: 		"S",
	icecream:	[]string{"chocolate", "vanilla"},
}

p2 := person{
	first: 		"Colista",
	last:		"S",
	icecream:	[]string{"strawberry", "vanilla"},
}

fmt.Println(p1)
fmt.Println(p2)

fmt.Println(p1.icecream)
fmt.Println(p2.icecream)

for _, v := range p1.icecream {
	fmt.Println(p1.first, "favorite is", v)
}

for _, v := range p2.icecream {
	fmt.Println(p2.first, "favorite is", v)
}
}