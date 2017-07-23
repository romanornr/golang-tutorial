package main

import (
    "fmt"
)

func rectProps(l, w float64)(float64, float64){
    area := l * w
    perimeter := (l+w) * 2
    return area, perimeter
}

func main(){
    //area, perimeter := rectProps(10.9, 3.2)
    //fmt.Printf("Area %f Perimeter %f", area, perimeter)
    area, _ := rectProps(10.9, 3.3)
    fmt.Printf("Area %f", area)
}
