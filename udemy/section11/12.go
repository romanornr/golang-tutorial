package main

import (
	"encoding/json"
	"os"
)

type Altcoin struct {
	Name   string
	Symbol string
	Price  float32
}

func main() {

	stratis := Altcoin{"Stratis", "STRAT", 80.22}
	json.NewEncoder(os.Stdout).Encode(&stratis)
}
