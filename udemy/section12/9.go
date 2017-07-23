package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println([]byte("Hello world"))
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(y)
	fmt.Println(y, z)
}
