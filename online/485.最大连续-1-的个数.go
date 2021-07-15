/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			continue
		}
		maxCount = max(maxCount, count)
		count = 0
	}
	maxCount = max(maxCount, count)

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
