package main

import (
	"encoding/json"
	"fmt"
)

type Altcoin struct {
	Name   string
	Symbol string
}

func main() {
	var alt Altcoin
	fmt.Println(alt.Name)

	bs := []byte(`{"Name":"Stratis", "Symbol":"STRAT"}`)
	json.Unmarshal(bs, &alt)

	fmt.Println(alt.Symbol)
}
