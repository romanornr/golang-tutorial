package main

import "fmt"

func main() {
    a, b := 5.4, 8.97
    fmt.Printf("type of a %T b %T\n", a, b)
    sum := a + b
    fmt.Println(sum)
}
