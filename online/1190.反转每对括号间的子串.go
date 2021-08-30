package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	stack := make([][]byte, 0)
	str := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, str)
			str = []byte{}
		} else if s[i] == ')' {
			for j := 0; j < len(str)/2; j++ {
				str[j], str[len(str)-j-1] = str[len(str)-j-1], str[j]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		} else {
			str = append(str, s[i])
		}
	}

	return string(str)
}

// @lc code=end

func main() {
	fmt.Println(reverseParentheses(`(u(love)i)`))
}

// 示例 1：
//
// 输入：s = "(abcd)"
// 输出："dcba"
// 示例 2：
//
// 输入：s = "(u(love)i)"
// 输出："iloveu"
// 解释：先反转子字符串 "love" ，然后反转整个字符串。
// 示例 3：
//
// 输入：s = "(ed(et(oc))el)"
// 输出："leetcode"
// 解释：先反转子字符串 "oc" ，接着反转 "etco" ，然后反转整个字符串。
// 示例 4：
//
// 输入：s = "a(bcdefghijkl(mno)p)q"
// 输出："apmnolkjihgfedcbq"
