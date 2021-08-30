package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	sort.SliceStable(nums, func(i, j int) bool {
		si := strconv.Itoa(nums[i])
		sj := strconv.Itoa(nums[j])

		ni, _ := strconv.Atoi(si + sj)
		nj, _ := strconv.Atoi(sj + si)
		return ni > nj
	})
	ans := ``
	for i := 0; i < len(nums); i++ {
		ans += strconv.Itoa(nums[i])
	}
	if nums[0] == 0 {
		return `0`
	}

	return ans
}

// @lc code=end

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
