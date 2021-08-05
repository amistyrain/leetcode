package main

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	// p := 0
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 0 {
	// 		nums[i], nums[p] = nums[p], nums[i]
	// 		p++
	// 	}
	// }
	// for i := p; i < len(nums); i++ {
	// 	if nums[i] == 1 {
	// 		nums[i], nums[p] = nums[p], nums[i]
	// 		p++
	// 	}
	// }
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}

// @lc code=end
