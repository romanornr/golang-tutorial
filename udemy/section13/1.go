package main

import (
	"fmt"
)

func test() {
	for i := 0; i < 40; i++ {
		fmt.Println("Test", i)
	}
}

func test2() {
	for i := 0; i < 40; i++ {
		fmt.Println("test2", i)
	}
}

func main() {
	go test()
	go test2()
}
