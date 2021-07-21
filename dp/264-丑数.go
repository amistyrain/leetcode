package main

import "fmt"

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		dp[i] = min(dp[p2]*2, min(dp[p3]*3, dp[p5]*5))
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}

	}

	return dp[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// @lc code=end

func main() {
	fmt.Println(nthUglyNumber(10))
}
