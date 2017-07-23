package main

import "fmt"

func main() {
	f := factorial(4)
	fmt.Println(f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= 1
	}
	return total
}
