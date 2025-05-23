package main

import (
	"fmt"
	"encoding/json"
)

type user struct {
	First   string		`json:"First"`
	Last    string		`json:"Last"`
	Age     int			`json:"Age"`
	Sayings []string	`json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you.","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var users []user

	err := json.Unmarshal(bs, &users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nAll of the data", users)

	for i, v := range users {
		fmt.Println("\n PERSON NUMBER", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, saying := range v.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}