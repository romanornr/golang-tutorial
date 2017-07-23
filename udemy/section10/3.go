package main

import (
	"fmt"
)

func main() {
	altcoins := map[string]uint16{
		"Litecoin": 2,
		"Stratis":  16,
	}

	altcoins["Crowncoin"] = 20
	fmt.Println(altcoins)

}
