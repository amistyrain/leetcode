package main

import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	ans := make([]string, n)
	for i := 1; i <= n; i++ {
		v := ""
		switch {
		case i%15 == 0:
			v = "FizzBuzz"
		case i%5 == 0:
			v = "Buzz"
		case i%3 == 0:
			v = "Fizz"
		default:
			v = strconv.Itoa(i)
		}
		ans[i-1] = v
	}

	return ans
}

// @lc code=end
