package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
