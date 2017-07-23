package main

import (
	"encoding/json"
	"fmt"
)

type Altcoin struct {
	Name   string  `json:"name"`
	Symbol string  `json:"sym"`
	Price  float32 `json:"price"`
}

func main() {
	alt := Altcoin{"Stratis", "STRAT", 30.20}
	bs, _ := json.Marshal(alt)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
