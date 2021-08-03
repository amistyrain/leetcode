/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	newNums := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			newNums = append(newNums, nums[i])
		}
	}

	if len(newNums) == 0 {
		return 1
	}
	for i := 0; i < len(newNums); i++ {
		num := abs(newNums[i])
		if num <= len(newNums) && newNums[num-1] > 0 {
			newNums[num-1] = -newNums[num-1]
		}
	}

	for i := 0; i < len(newNums); i++ {
		if newNums[i] > 0 {
			return i + 1
		}
	}

	return len(newNums) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// @lc code=end

