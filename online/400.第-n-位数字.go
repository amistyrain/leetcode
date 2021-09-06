package main

/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第 N 位数字
 */

// @lc code=start
func findNthDigit(n int) int {
	base := 9
	digits := 1
	for n-digits*base > 0 {
		n -= digits * base
		base *= 10
		digits++
	}
	idx := n % digits
	if idx == 0 {
		idx = digits
	}
	num := 1
	for i := 1; i < digits; i++ {
		num *= 10
	}
	if idx == digits {
		num += n/digits - 1
	} else {
		num += n / digits
	}
	for i := idx; i < digits; i++ {
		num /= 10
	}

	return num % 10
}

// @lc code=end
