package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	absoluteZero Celsius = -273.15
	freezing     Celsius = 0
	boling       Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func main() {
	fmt.Printf("%g\n", boling-freezing)
	fmt.Printf("%g\n", boling-CToF(freezing))
}
