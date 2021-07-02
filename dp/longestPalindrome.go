package main

import "fmt"

func longestPalindrome(str string) string {
	res := ``

	for i := 0; i < len(str); i++ {
		s1 := palindrome(&str, i, i)
		s2 := palindrome(&str, i, i+1)
		res = max(res, s1)
		res = max(res, s2)
	}

	return res
}

func palindrome(str *string, l, r int) string {
	for l >= 0 && r < len(*str) && (*str)[l] == (*str)[r] {
		l--
		r++
	}
	if l+1 > r {
		return (*str)[l:r]
	}
	return (*str)[l+1 : r]
}

func max(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}

	return s2
}

func main() {
	fmt.Println(longestPalindrome(`mmlabcdcbahjkk`))
	fmt.Println(longestPalindrome(`mmlabcddcbahjkk`))
	fmt.Println(longestPalindrome(`mm`))
}
