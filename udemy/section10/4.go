package main

import (
	"fmt"
	"sort"
)

func main() {
	altcoins := []string{
		1: "Litecoin",
		2: "Stratis",
		3: "Decred",
	}
	sort.Ints(altcoins.key)

	for key, value := range altcoins {
		fmt.Println(key, "-", value)
	}
}
