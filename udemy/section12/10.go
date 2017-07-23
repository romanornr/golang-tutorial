package main

import "fmt"

func main() {
	var name interface{} = "Altcoin"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		println("Not a value")
	}

	var val interface{} = 7
	fmt.Println(val.(int) + 6)
}
