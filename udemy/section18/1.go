package main

import (
	"os"
)

func main() {
	_, err := os.Open("test.txt")
	if err != nil {
		//fmt.Println("error: ", err)
		//1log.Println(err)
		//log.Fatalln(err)
		panic(err)
	}
}
