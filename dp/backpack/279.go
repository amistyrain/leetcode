package main

import "fmt"

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	coins := []int{}
	for i := 1; i*i <= n; i++ {
		coins = append(coins, i*i)
	}
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	fmt.Println(numSquares(13))
}
