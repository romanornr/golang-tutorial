package main

import (
    "fmt"
    "math"
)

func main() {
    a, b := 145.8, 10.0
    c := math.Min(a, b)
    fmt.Println("min value is ", c)
}
