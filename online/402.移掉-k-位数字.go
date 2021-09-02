package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) != 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]

	n := strings.TrimLeft(string(stack), "0")
	if len(n) == 0 {
		return `0`
	}

	return n
}

// @lc code=end

// 输入：num = "1432219", k = 3
// 输出："1219"
// 解释：移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219 。

func main() {
	fmt.Println(removeKdigits(`1432219`, 3))
	fmt.Println(removeKdigits(`10200`, 1))
}
