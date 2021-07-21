package main

import "fmt"

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextElemets := make(map[int]int, len(nums2))
	stack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nextElemets[nums2[i]] = -1
		} else {
			nextElemets[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = nextElemets[nums1[i]]
	}

	return result
}

// @lc code=end

func main() {

	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}

// 输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
// 输出: [-1,3,-1]
// 解释:
//     对于 num1 中的数字 4 ，你无法在第二个数组中找到下一个更大的数字，因此输出 -1 。
//     对于 num1 中的数字 1 ，第二个数组中数字1右边的下一个较大数字是 3 。
//     对于 num1 中的数字 2 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
