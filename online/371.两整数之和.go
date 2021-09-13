package main

/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	for a != 0 {
		c := a ^ b
		d := (a & b) << 1
		a = c
		b = d
	}

	return b
}

// @lc code=end
