package main

import (
	"fmt"
)

func main() {
	n := HashBucket("go", 12)
	fmt.Println(n)
}

func HashBucket(word string, buckets int32) int32 {
	letter := rune(word[0])
	bucket := letter % buckets
	return bucket
}
