package main

import (
	"fmt"
	"sort"
)

type altcoins []string

func (a *altcoins) Len() int          { return len(*a) }
func (a altcoins) Less(i, j int) bool { return a[i] < a[j] }
func (a altcoins) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	s := altcoins{"Litecoin", "Decred", "Zcoin", "Stratis"}
	sort.Sort(&s)
	fmt.Println(s)
}
