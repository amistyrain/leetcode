package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	if len(S) == 0 {
		return ``
	}
	j := 0
	char := S[0]
	count := 1
	s := ``
	for i := 1; i < len(S); i++ {
		if S[i] == char {
			count++
			continue
		}
		s += string(char) + strconv.Itoa(count)
		char = S[i]
		count = 1
		j += 2
	}
	s += string(char) + strconv.Itoa(count)
	if len(string(s)) >= len(S) {
		return S
	}

	return string(s)
}

func main() {
	fmt.Println(compressString(`aabcccccaaa`))
	fmt.Println(compressString(`abbd`))
}
