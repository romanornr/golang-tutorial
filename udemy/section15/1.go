package main

import (
	"fmt"
)

func main() {
	c1 := incrementor("test1: ")
	c2 := incrementor("test2: ")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("final counter: ", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
