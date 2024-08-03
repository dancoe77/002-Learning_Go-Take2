package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("This is where initializiation for my program occurs")
}

func main(){
	x := rand.Intn(250)
	fmt.Printf("x is %v \n", x)

	switch {
	case x <=100:
		fmt.Println("x is between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x >=201 && x <= 250:
		fmt.Println("x is between 201 and 250")
	}
	/*
	The init function¶
	Finally, each source file can define its own niladic init function to set up whatever state is required. (Actually each file can have multiple init functions.) And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.
	
	Besides initializations that cannot be expressed as declarations, a common use of init functions is to verify or repair correctness of the program state before real execution begins.
	
	func init() {
	    if user == "" {
	        log.Fatal("$USER not set")
	    }
	    if home == "" {
	        home = "/home/" + user
	    }
	    if gopath == "" {
	        gopath = home + "/go"
	    }
	    // gopath may be overridden by --gopath flag on command line.
	    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
	}
	[The init function](https://go.dev/doc/effective_go#init)
	
	niladic

	(computing) Of an operator or function in a program, having no arguments.
	
	[niladic](https://en.wiktionary.org/wiki/niladic)
	*/
}	