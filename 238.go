package main

import "fmt"

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	left := 1
	for i := 0; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}
