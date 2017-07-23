package main

import (
	"fmt"
)

func main() {

	records := make([][]string, 0)
	altcoin1 := make([]string, 5)
	altcoin1[0] = "Litecoin"
	altcoin1[1] = "0.3"
	altcoin1[2] = "0.1"
	altcoin1[3] = "4.5% change"
	altcoin1[4] = "2280 BTC volume"

	altcoin2 := make([]string, 5)
	altcoin2[0] = "Stratis"
	altcoin2[1] = "0.2"
	altcoin2[2] = "0.25"
	altcoin2[3] = "56% change"
	altcoin2[4] = "31412 BTC volume"

	test := append(records, altcoin1, altcoin2)
	fmt.Println(test)
}
