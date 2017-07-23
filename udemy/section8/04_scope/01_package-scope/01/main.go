package main

import (
	"fmt"
)

var x int = 42 //package level scope

//package level scopes are accesible inside scopes that
// are contained in the package level scopes.

//so the package level is the outer scope, enclosing to other scopes
// which are package level scopes

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
