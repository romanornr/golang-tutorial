package main

import (
	"fmt"
)

func main() {
	var altcoin = make(map[string]string)
	altcoin["Litecoin"] = "Poloniex"

	fmt.Println(altcoin)
}
