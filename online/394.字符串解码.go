package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

// @lc code=end

// 示例 1：
//
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
// 示例 2：
//
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
// 示例 3：
//
// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
// 示例 4：
//
// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"

func main() {
	fmt.Println(decodeString(`abc`))
}
