package main

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	k := 0
	for i := 0; i <= k; i++ {
		k = max(k, i+nums[i])
		if k >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
func main() {

}
