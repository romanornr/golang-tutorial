package main

import (
	"fmt"
)

type altcoin struct {
	name   string
	symbol string
	price  float32
}

func main() {
	stratis := altcoin{"Stratis", "STRAT", 12.20}
	fmt.Println(stratis.name, stratis.price)
}
