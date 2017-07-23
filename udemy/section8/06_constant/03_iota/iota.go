package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n", kb)
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\n", mb)
}
