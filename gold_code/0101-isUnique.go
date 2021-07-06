package main

import "fmt"

func isUnique(astr string) bool {
	count := 0

	for i := 0; i < len(astr); i++ {
		bit := rune(astr[i]) - 'a'
		if count&(1<<bit) != 0 {
			return false
		}
		count |= 1 << bit
	}

	return true
}

func main() {
	fmt.Println(isUnique(`leetcode`))
}
