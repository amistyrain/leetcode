/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	k := 1
	for res&k == 0 {
		k <<= 1
	}
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&k == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	return []int{a, b}
}

// @lc code=end

