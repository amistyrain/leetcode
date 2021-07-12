package main

import "fmt"

func canPermutePalindrome(s string) bool {
	bits := [4]uint64{}
	for i := 0; i < len(s); i++ {
		num := rune(s[i])
		bitIndex := num / 64
		mvNum := num % 64
		fmt.Println(mvNum)
		bits[bitIndex] = bits[bitIndex] ^ (1 << mvNum)
	}
	res := uint64(0)
	for i := 0; i < len(bits); i++ {
		res ^= bits[i]
	}
	if res == 0 || res&(res-1) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(canPermutePalindrome(`tactcoa`))
	fmt.Println(canPermutePalindrome(`tactcooa`))
	fmt.Println(canPermutePalindrome(`tactcb`))
	fmt.Println(canPermutePalindrome(`AaBb//a`))
	fmt.Println(canPermutePalindrome(`Aac`))
	fmt.Println(canPermutePalindrome(`\a`))
}
