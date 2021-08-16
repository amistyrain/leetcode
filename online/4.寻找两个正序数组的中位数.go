package main

import "fmt"

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	l1, l2 := 0, 0
	num1, num2 := -1, -1
	for i := 0; i <= (len1+len2)/2; i++ {
		num1 = num2
		if l1 < len1 && (l2 >= len2 || nums1[l1] < nums2[l2]) {
			num2 = nums1[l1]
			l1++
		} else {
			num2 = nums2[l2]
			l2++
		}
	}

	if (len1+len2)%2 == 1 {
		return float64(num2)
	}

	return float64(num1+num2) / 2.0
}

// @lc code=end

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
	fmt.Println(findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
}
