package main

import "fmt"

func main() {
	greeting := []string{
		"hi",
		"Good morning",
		"Nice weather",
		"pls staph",
		"snek",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
