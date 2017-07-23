package main

import "fmt"

func main() {
	var x = 2
	var y = 12.12
	fmt.Println(y + float64(x))

	var z rune = 'a'
	fmt.Println(string(z))
}
