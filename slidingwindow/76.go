package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	start := 0
	minLen := math.MaxInt64
	windows := make(map[byte]int, len(t))
	needs := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}
	left, right := 0, 0
	match := 0
	for right < len(s) {
		char := s[right]
		if _, ok := needs[char]; ok {
			windows[char]++
			if windows[char] == needs[char] {
				match += 1
			}
		}
		right++

		for match == len(needs) {
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			char = s[left]
			if _, ok := needs[char]; ok {
				windows[char]--
				if windows[char] < needs[char] {
					match--
				}
			}
			left++
		}

	}
	if minLen == math.MaxInt64 {
		return ``
	}

	return s[start : start+minLen]
}

func main() {
	// fmt.Println(minWindow(`ADOBECODEBANC`, `ABC`))
	fmt.Println(minWindow(`aa`, `aa`))
}
