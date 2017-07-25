package main

import (
	"fmt"
	"log"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("A norgate marh error: %v &v %v\n", n.lat, n.long, n.err)

}

func main() {
	_, err := Sqrt(-10.20)
	if err != nil {
		log.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v\n", f)
		return 0, &NorgateMathError{"50.2289 N", "99.34 W", nme}
	}
	return 42, nil
}
