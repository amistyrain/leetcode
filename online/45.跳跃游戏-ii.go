package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	maxPosition := 0
	end := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			step++
			end = maxPosition
		}
	}

	return step
}

// dp
// func jump(nums []int) int {
// 	dp := make([]int, len(nums))
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = math.MaxInt64
// 	}
// 	dp[0] = 0
//
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums) && j <= i+nums[i]; j++ {
// 			dp[j] = min(dp[i]+1, dp[j])
// 		}
// 	}
//
// 	return dp[len(nums)-1]
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
