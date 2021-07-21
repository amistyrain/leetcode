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

