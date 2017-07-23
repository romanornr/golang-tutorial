package main

import "fmt"

type altcoin struct {
	name   string
	symbol string
	price  float64
}

func (a altcoin) getDollarPrice() float64 {
	return a.price * 2500
}

func main() {
	a := altcoin{"Stratis", "STRAT", 0.123}
	price := a.getDollarPrice()
	fmt.Printf("$%.2f\n", price)
}
