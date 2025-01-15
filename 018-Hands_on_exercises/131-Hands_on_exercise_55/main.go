package main

import "fmt"

type engine struct {
	electric bool
	}

type vehicle struct {
	engine
	manufacturer 	string
	model 			string
	doors			int
	color			string
	}

func main (){
	ve1 := vehicle {
		engine: engine{
			electric: true,
		},
		manufacturer:	"Toyota",
		model:			"Prius",
		doors:			4,
		color:			"blue",
	}

	ve2 := vehicle {
		engine: engine{
			electric: true,
		},
		manufacturer:	"Aptera",
		model:			"Production Intent",
		doors:			2,
		color:			"black",
	}

	fmt.Println(ve1)
	fmt.Println(ve2)

	fmt.Printf("%T \t %#v\n", ve1, ve1)
	fmt.Printf("%T \t %#v\n", ve2, ve2)

	fmt.Println(ve1.manufacturer, ve1.model, ve1.doors, ve1.color, ve1.engine.electric)
	fmt.Println(ve2.manufacturer, ve2.model, ve2.doors, ve2.color, ve2.engine.electric)
}