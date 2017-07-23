package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {

	x, even := half(8)
	fmt.Println(x, even)
}
