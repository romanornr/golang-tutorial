package main

import (
	"fmt"
)

func main() {
	greeting := make([]string, 3, 3)

	greeting[0] = "hi"
	greeting[1] = "Two"
	greeting[2] = "wow"

	greeting = append(greeting, "hihi")

	fmt.Println(greeting)
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
