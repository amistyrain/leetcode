package main

import "fmt"

func replaceSpaces(S string, length int) string {
	if length == 0 {
		return ``
	}
	spaceCount := 0
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			spaceCount++
		}
	}
	fmt.Println(spaceCount)
	newLength := length + 2*spaceCount
	newS := make([]byte, newLength)

	for i, j := length-1, newLength-1; i >= 0 && j >= 0; {
		if S[i] == ' ' {
			newS[j] = '0'
			newS[j-1] = '2'
			newS[j-2] = '%'
			j = j - 2
		} else {
			newS[j] = S[i]
		}
		i--
		j--
	}

	return string(newS)
}

func main() {
	fmt.Println(replaceSpaces(`Mr John Smith`, 13))
	fmt.Println(replaceSpaces(`     `, 5))
}
