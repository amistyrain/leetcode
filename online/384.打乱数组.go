package main

import "math/rand"

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	old []int
}

func Constructor(nums []int) Solution {
	return Solution{old: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.old
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, 0, len(this.old))
	nums = append(nums, this.old...)
	n := len(nums)
	x := n
	for i := 0; i < n-1; i++ {
		randIdx := rand.Intn(x)
		nums[i], nums[i+randIdx] = nums[i+randIdx], nums[i]
		x--
	}

	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end
