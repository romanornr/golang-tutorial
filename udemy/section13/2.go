package main

import "fmt"
import "sync"
import "time"
import "runtime"

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func test3() {
	for i := 0; i < 44; i++ {
		fmt.Println("test", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func test4() {
	for i := 0; i < 42; i++ {
		fmt.Println("Test2", i)
		time.Sleep(time.Duration(3 + time.Millisecond))
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go test3()
	go test4()
	wg.Wait()
}
