package main
import "fmt"

func main() {
    var a bool = true
    b := false
    fmt.Println(a, b)
    c := a && b
    fmt.Println(c)
    d := a || b
    fmt.Println(d)
}
