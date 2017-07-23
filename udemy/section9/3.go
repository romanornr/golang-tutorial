package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 0, 3)
	fmt.Println(len(mySlice))
	fmt.Println("----")
	fmt.Println(cap(mySlice))

	for i := 0; i < 87; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("len: %d, cap: %d\n", len(mySlice), cap(mySlice))
	}

}
