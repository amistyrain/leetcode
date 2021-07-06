package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	bit := 0
	s := s1 + s2
	sums1, sums2 := 0, 0
	for i := 0; i < len(s1); i++ {
		sums1 += int(s1[i])
		sums2 += int(s2[i])
	}
	for i := 0; i < len(s); i++ {
		bit ^= int(s[i])
	}

	return bit == 0 && sums1 == sums2
}

func main() {
	fmt.Println(CheckPermutation(`abc`, `acb`))
	fmt.Println(CheckPermutation(`abc`, `acd`))
	fmt.Println(CheckPermutation(`aa`, `bb`))
}
