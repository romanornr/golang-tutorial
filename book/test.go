package main

import (
	"fmt"
)

func eq(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	x := "hi"
	y := [2]string{"hi", "no"}
	fmt.Println(eq(x, y))
}
