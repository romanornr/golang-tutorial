package main

import "fmt"

func main() {
	age := 44
	fmt.Println(age)

	ChangeMe(&age)

}

func ChangeMe(x *int) { //take a type, pointer to an int
	fmt.Println(x)  //memory addr
	fmt.Println(*x) //44
	*x = 24
	fmt.Println(x)  //same memory addr
	fmt.Println(*x) // 24
}
