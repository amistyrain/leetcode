package main

import "fmt"

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}

func main() {
	fmt.Println(firstUniqChar(`abaccdeff`))
	fmt.Println(firstUniqChar(``))

}
