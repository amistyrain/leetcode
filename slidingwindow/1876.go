package main

import "fmt"

func countGoodSubstrings(s string) int {
	res := 0
	windows := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		char := s[right]
		windows[char]++
		right++
		for windows[char] > 1 {
			char2 := s[left]
			windows[char2]--
			left++
		}
		if right-left == 3 {
			res++
			char2 := s[left]
			windows[char2]--
			left++
		}
	}

	return res
}

func main() {
	fmt.Println(countGoodSubstrings(`aababcabc`))
	fmt.Println(countGoodSubstrings(`aababcabc`))
}
