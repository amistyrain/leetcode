package main

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	ans := 0
	sum := make(map[int]int, len(nums1)*2)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			sum[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			if _, ok := sum[-n3-n4]; ok {
				ans += sum[-n3-n4]
			}
		}
	}

	return ans
}

// @lc code=end
