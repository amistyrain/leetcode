package main

import "fmt"

func trailingZeroes(n int) int {
	res := 0
	i := 5

	for i <= n {
		res += n / i
		i *= 5
	}

	return res
}

func main() {
	fmt.Println(trailingZeroes(125))
}
