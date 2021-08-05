package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	maps := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = true
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		if !maps[nums[i]-1] {
			count := 1
			n := nums[i] + 1
			for maps[n] {
				count++
				n++
			}
			result = max(result, count)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
