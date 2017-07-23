package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)  // memory addr
	fmt.Println(*b) // 43

	*b = 42        //reasign value
	fmt.Println(a) // 42
}
