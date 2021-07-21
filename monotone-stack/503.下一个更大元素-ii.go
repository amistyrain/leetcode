package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	stack := []int{}
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i%n] = -1
		} else {
			result[i%n] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%n])
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
