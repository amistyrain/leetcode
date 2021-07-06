package main

import "fmt"

func convertInteger(A int, B int) int {
	count := 0
	num := int32(A ^ B)

	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func main() {
	fmt.Println(convertInteger(29, 15))
	fmt.Println(convertInteger(1, 2))
}
