package main

import (
	"fmt"
)

type Bitcoin struct {
	price     float64
	algorithm string
}

type Altcoin struct {
	Bitcoin
	name string
	dev  string
}

func main() {
	altcoin1 := Altcoin{
		Bitcoin: Bitcoin{
			price:     30.1,
			algorithm: "Scrypt",
		},
		name: "Stratis",
		dev:  "Stratis devs",
	}

	fmt.Println(altcoin1.algorithm)
}
