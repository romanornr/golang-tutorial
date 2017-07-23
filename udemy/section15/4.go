package main

import "fmt"

func main() {
	f := factorial(4)
	for n := range f {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	c := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
