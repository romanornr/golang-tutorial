package main

import "fmt"

type Vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type Car struct {
	Vehicle
	Wheels int
	Doors  int
}

type plane struct {
	Vehicle
	Jet bool
}

type boat struct {
	Vehicle
	Length int
}

func (v Vehicle) Specs() {
	fmt.Printf("seats %v, max speed %v, color %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	prius := Car{}
	tacoma := Car{}
	bmw528 := Car{}
	cars := []Car{prius, tacoma, bmw528}
	fmt.Println(cars)
}
