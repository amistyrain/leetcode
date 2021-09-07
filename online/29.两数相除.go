package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend <= math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	sign := 1
	// 异号
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		sign = -1
	}
	a := dividend
	b := divisor
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	ans := div(a, b)
	if sign < 0 {
		return -ans
	}

	return ans
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	tb := b
	count := 1
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}

	return count + div(a-tb, b)
}

// @lc code=end

// 10 3
// 1010
// 0011

// 7  3
// 0111
// 0011
