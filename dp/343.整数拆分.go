package main

import "fmt"

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i/2+1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), dp[i-j]*j))
		}

	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// @lc code=end

func main() {
	// 3 3 4  3*3*4=36
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak(2))
}
