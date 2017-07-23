package main

import (
	"fmt"
	"sort"
)

type altcoins []string

func main() {
	s := []string{"Litecoin", "Decred", "Zcoin", "Stratis"}
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

}
