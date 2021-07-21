package main

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if guess(mid) <= 0 {
			right = mid
		} else {
			left = mid + 1
		}

	}

	return left
}

// @lc code=end
