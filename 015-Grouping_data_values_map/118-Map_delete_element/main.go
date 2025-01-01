package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
		}

		fmt.Println("The age of Henry was", am["Henry"])

		fmt.Println(am)
		fmt.Printf("%#v\n", am)

		an := make(map[string]int)
		an["Lucas"] = 28
		an["Steph"] = 37
		an["George"] = 78
		fmt.Println(an)
		fmt.Printf("%#v\n", an)
		fmt.Println(len(an))

		delete (an, "George")
		//delete (an, "Georage")//won't panic when run
		fmt.Println(an)
		fmt.Printf("%#v\n", an)
		fmt.Println(len(an))

		//fmt.Println("--- accessing keys that don't exist")
		//fmt.Println(an["Georgey"]) = won't panic either
		//fmt.Println("--------------------------------------------")

		// for range over a MAP
		for k, v := range an {
			fmt.Println(k, v)
		}

		for _, v := range an {
			fmt.Println(v)
		}

		for k := range an {
			fmt.Println(k)
		}

		//for range with a SLICE

		xi := []int{42, 43, 44}

		for i, v := range xi {
			fmt.Println(i, v)
		}

		for _, v := range xi {
			fmt.Println(v)
		}

		for i := range xi {
			fmt.Println(i)
		}




		/*
			m := map[string]int{
				"todd":  42,
				"henry": 16,
			}

			//  --- or ---

			// m := make(map[string]int)
			// m["todd"] = 42
			// m["henry"] = 16

			fmt.Println("Henry's age is ", m["henry"])

			for k := range m {
				fmt.Printf("just the keys: %s\n", k)
			}

			for k, v := range m {
				fmt.Printf("%s is %d years old\n", k, v)
			}

			for _, v := range m {
				fmt.Printf("just the values: %d\n", v)
			}

			if v, ok := m["Padget"]; ok {
				fmt.Printf("%s is %d years old\n", "Padget", v)
			} else {
				fmt.Printf("%s not found\n", "Padget")
			}

			//delete
			m["Shakespeare"] = 459
			fmt.Printf("%#v\n", m)
			delete(m, "Shakespeare")
			fmt.Printf("%#v\n", m)
			delete(m, "Shakespeare") // no panic

			// len
			fmt.Println("len: ", len(m))
		*/
	}

	//  --- or ---

	/*
	var m map[string]int
	fmt.Println(m["tunde"])
	m = make(map[string]int)
	m["todd"] = 42
	m["henry"] = 16
	*/
