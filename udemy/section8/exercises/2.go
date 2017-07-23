package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	biggest := max(4, 5, 9, 20, 17)
	fmt.Println(biggest)
}
