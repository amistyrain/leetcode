package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	nums := make([]int, 0, len(candidates))
	backTrack(candidates, 0, nums, target)

	return res
}

var res [][]int

func backTrack(candidates []int, start int, nums []int, target int) {
	if target == 0 {
		res = append(res, append([]int{}, nums...))
		return
	}
	if target < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		nums = append(nums, candidates[i])

		backTrack(candidates, i+1, nums, target-candidates[i])

		nums = nums[:len(nums)-1]
	}
}

// @lc code=end

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
