package main

var pc [256]byte

func init() {
	for i:= range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
	pc[]
}