package main

import "fmt"

type price float32

func main() {
	var Litecoin price
	Litecoin = 44
	fmt.Printf("%T %v \n", Litecoin, Litecoin)
}
