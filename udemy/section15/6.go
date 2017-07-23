package main

func main() {
	c := gen(100)
	f := factorial(c)
}

func gen(number int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < number; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(n chan int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
