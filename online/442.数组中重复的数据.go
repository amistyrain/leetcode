/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			res = append(res, abs(nums[i]))
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

