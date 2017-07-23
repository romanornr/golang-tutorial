package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)  // 43
	fmt.Println(&a) // memory addr

	var b *int = &a //referencing a memory addr where an int is stored
	fmt.Println(b)  // memory addr
	fmt.Println(*b) // 43 star means derefrence
}
