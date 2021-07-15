package main

import "fmt"

/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= len(coins); i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}

// @lc code=end

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
}
