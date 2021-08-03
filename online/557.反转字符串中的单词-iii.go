package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	bs := []byte(s)

	for l, r := 0, len(bs)-1; l < r; l, r = l+1, r-1 {
		bs[l], bs[r] = bs[r], bs[l]
	}

	s = string(bs)
	strs := strings.Split(string(bs), ` `)

	for l, r := 0, len(strs)-1; l < r; l, r = l+1, r-1 {
		strs[l], strs[r] = strs[r], strs[l]
	}

	return strings.Join(strs, ` `)
}

// @lc code=end

func main() {
	fmt.Println(reverseWords(`Let's take LeetCode contest`))

}
