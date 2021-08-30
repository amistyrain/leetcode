package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.SliceStable(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			(envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})

	heights := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		heights[i] = envelopes[i][1]
	}

	return lengthOfLIS(heights)
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	fmt.Println(maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
}
