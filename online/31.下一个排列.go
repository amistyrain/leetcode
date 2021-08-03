package main

import "fmt"

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
			fmt.Println()
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]

	}

	return
}

// @lc code=end

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
