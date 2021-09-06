package main

import "fmt"

/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start
func maxCoins(nums []int) int {
	n := len(nums)
	store := make([]int, n+2)
	store[0] = 1
	store[len(store)-1] = 1
	for i := 0; i < n; i++ {
		store[i+1] = nums[i]
	}
	dp := make([][]int, n+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+2)
	}
	//dp[i][j] = dp[i][k] + val[i]*val[k]*val[j] + dp[k][j]
	//L 区间大小
	for l := 3; l <= n+2; l++ {
		for i := 0; i <= n+2-l; i++ {
			res := 0
			for k := i + 1; k < i+l-1; k++ {
				left := dp[i][k]
				right := dp[k][i+l-1]
				res = max(res, left+store[i]*store[k]*store[i+l-1]+right)
			}
			dp[i][i+l-1] = res
		}
	}
	fmt.Println(dp)

	return dp[0][n+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
