package main

import "fmt"

const (
	A = iota
	B = iota * 10
	C = iota * 10
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
