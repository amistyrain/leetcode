/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			return abs(nums[i])
		}
		if nums[index] > 0 {
			nums[index] *= -1
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

// @lc code=end

