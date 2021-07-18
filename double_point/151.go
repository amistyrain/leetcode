package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	result := ``
	left, right := len(s)-1, len(s)-1
	for left >= 0 {
		for left == right && left >= 0 && s[left] == ' ' {
			left--
			right--
		}
		for left >= 0 && s[left] != ' ' {
			left--
		}
		result += s[left+1:right+1] + ` `
		right = left
	}

	return result[0 : len(result)-1]
}

// @lc code=end

func main() {
	fmt.Println(reverseWords(`  Bob    Loves  Alice   `))
	fmt.Println(reverseWords(`a good   example`))
	fmt.Println(reverseWords(` asdasd df f`))
	fmt.Println(reverseWords(`  asdasd df f`))
	fmt.Println(reverseWords(`asdasd df f`))
	fmt.Println(reverseWords(`  hello world  `))
	fmt.Println(reverseWords(`  `))
}
