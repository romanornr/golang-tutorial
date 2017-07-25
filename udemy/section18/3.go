package main

import "errors"
import "log"

var ErrNorGateMath = errors.New("norgate Math: square root of negative")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(ErrNorGateMath)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorGateMath
	}
	return 42, nil
}
