package main

import "fmt"

const p string = "death & taxes"

func main() {
	const q = 42

	fmt.Printf("p - %s\n", p)
	fmt.Printf("q - %d\n", q)
}
