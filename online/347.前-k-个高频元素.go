package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	result := make([]int, 0, k)
	arr := make([][]int, 0, len(count))
	for k, v := range count {
		arr = append(arr, []int{v, k})
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})

	for i := 0; i < len(arr); i++ {
		if i < k {
			result = append(result, arr[i][1])
			continue
		}
		break
	}

	return result
}

// 最小堆
// @lc code=end

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
}
