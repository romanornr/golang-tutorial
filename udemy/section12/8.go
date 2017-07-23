package main

import "fmt"

type animal struct {
	sound string
}

type mouse struct {
	animal
	friendly bool
}

type dog struct {
	animal
	friendly bool
}

func main() {
	test1 := dog{animal{"woof"}, true}
	test2 := mouse{animal{"..."}, false}
	test3 := []interface{}{test1, test2}
	fmt.Println(test3)
}
