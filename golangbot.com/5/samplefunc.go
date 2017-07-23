package main

import(
    "fmt"
)

func calculateBill(price int, no int) int{
    totalprice := price * no
    return totalprice
}

func main(){
    fmt.Println(calculateBill(10, 2))
}
