package main

import (
	"fmt"
)

// 5 + 3
// 000
// 011

// 100

// 101
// 111

func add(a int, b int) int {
	for b != 0 {
		c := (a & b) << 1
		a = a ^ b
		b = c
	}

	return a
}

func main() {
	fmt.Println(add(5, 3))
	fmt.Println(5 & 3)
	fmt.Println(4 & 3)
}
