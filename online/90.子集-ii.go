package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	ans = make([][]int, 0)
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	track := make([]int, 0, len(nums))
	backTrack(nums, 0, track)

	return ans
}

var ans [][]int

func backTrack(nums []int, start int, track []int) {
	ans = append(ans, append([]int(nil), track...))

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backTrack(nums, i+1, track)

		track = track[:len(track)-1]
	}

}

// @lc code=end

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1, 1, 2, 2}))
}
