package main

import "fmt"

func main() {
	b := []byte{0, 1, 2, 3, 4, 5, 6}
	var v uint64 = 0x0807060504030201

	_ = b[7] // early bounds check to guarantee safety of writes below

	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)

	fmt.Println(b)
}