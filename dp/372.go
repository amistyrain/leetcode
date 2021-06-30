package main

import "fmt"

// 双指针
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := len(s) - 1
	j := len(t) - 1

	for j >= 0 {
		if s[i] == t[j] {
			if i == 0 {
				return true
			}
			j--
			i--
			continue
		}
		j--
	}

	return false
}

func main() {
	fmt.Println(isSubsequence(`b`, `abc`))
}
