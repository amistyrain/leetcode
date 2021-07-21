package main

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if 0 == x {
		return 0
	}

	n := x
	for n*n > x {
		n = (n + x/n) / 2
	}
	return n
}

func musqrt(x int) int {
	left, right := 0, x
	result := 0
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// @lc code=end
