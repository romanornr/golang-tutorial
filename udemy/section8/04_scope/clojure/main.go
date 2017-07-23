package main

import (
	"fmt"
)

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func test() func() int {
	x := 0
	return func() int {
		x = +8
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	test2 := test()
	fmt.Println(test2())
	fmt.Println(test2())
}
