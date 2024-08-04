package main

import (
	"fmt"
)

func main (){
	var bond = map[string]int{
		"James":		42,
		"Moneypenny":	32,
		"Q":			22,
	}
	for k, v := range bond {
		fmt.Println("ranging over a map", k, v)
	}
	m1 := bond["James"]
	fmt.Println(m1)
	
	m2 := bond["Q"]
	fmt.Println(m2)
	
	func m2(Q string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
	}
}