package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&counter, 2)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	incrementor("test12")
	incrementor("test14")
	wg.Wait()
}
