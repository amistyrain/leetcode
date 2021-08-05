package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Permutation([]byte{'0', '1', '2'}, 2))
}

func Permutation(seed []byte, size int) []string {
	result := backTrack(seed, []byte{}, size)
	sort.Strings(result)
	return result
}

func exist(seed []byte, c byte) bool {
	for i := 0; i < len(seed); i++ {
		if seed[i] == c {
			return true
		}
	}

	return false
}

func backTrack(seed []byte, chars []byte, size int) []string {
	if size == 0 {
		return []string{string(chars)}
	}

	result := make([]string, 0)
	for i := 0; i < len(seed); i++ {
		if exist(chars, seed[i]) {
			continue
		}
		chars = append(chars, seed[i])
		strs := backTrack(seed, chars, size-1)
		for j := 0; j < len(strs); j++ {
			result = append(result, strs[j])
		}

		chars = chars[:len(chars)-1]
	}

	return result
}
