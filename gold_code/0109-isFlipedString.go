package main

import "fmt"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	s := s1 + s1

	for i := 0; i < len(s1); i++ {
		if s[i:i+len(s1)] == s2 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(isFlipedString(`waterbottle`, `erbottlewat`))
	fmt.Println(isFlipedString(`aa`, `aba`))
}
