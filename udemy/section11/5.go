package main

import "fmt"

type Altcoin struct {
	Name   string
	Symbol string
}

type Working struct {
	Altcoin
	tradeable bool
}

func (a Altcoin) Exchange() {
	fmt.Println("This altcoin is on Bittrex")
}

func main() {
	a1 := Altcoin{
		Name:   "Stratis",
		Symbol: "STRAT",
	}

	p1 := Working{
		Altcoin: Altcoin{
			Name:   "Litecoin",
			Symbol: "LTC",
		},
		tradeable: true,
	}

	a1.Exchange()
	p1.Exchange()
	fmt.Println(p1.Symbol)

}
