package main

import "strings"
import "encoding/json"
import "fmt"

type Altcoin struct {
	Name   string
	Symbol string
}

func main() {
	var stratis Altcoin
	r := strings.NewReader(`{"Name":"Stratis", "symbol":"STRAT"}`)
	json.NewDecoder(r).Decode(&stratis)

	fmt.Println(stratis.Symbol)
}
