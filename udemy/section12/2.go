package main

import (
	"fmt"
	"sort"
)

func main() {

	type people []string
	//studyGroup := people{"Zeno", "John", "Jenny"}
	s := []string{"Zeno", "John", "Jenny"}

	n := []int{7, 4, 8, 23, 9, 19, 12, 32, 3}

	//sort.Sort(sort.StringSlice(s))
	sort.Strings(s)
	sort.Sort(sort.IntSlice(n))
	//fmt.Println(a)
	fmt.Println(s)
	fmt.Println(n)
}
