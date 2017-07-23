package main

import (
	"fmt"
)

type Altcoin struct {
	name   string
	symbol string
}

func main() {

	stratis := &Altcoin{"Stratis", "STRAT"}
	fmt.Println(stratis)
	fmt.Printf("%T\n", stratis)
	fmt.Println(stratis.symbol)
}
