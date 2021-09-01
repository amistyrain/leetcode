package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	if m < 2 {
		return total
	}

	return backTracking(nums, 0, m-1, total, math.MaxInt64)
}

func backTracking(nums []int, start int, splitNum int, sum int, splitSum int) int {
	minNum := math.MaxInt64
	if splitNum == 1 {
		n := 0
		for i := start; i < len(nums); i++ {
			n += nums[i]
			minNum = min(minNum, max(n, sum-n))
		}
		return min(splitSum, minNum)
	}
	if splitSum == math.MaxInt64 {
		splitSum = 0
	}

	for i := start; i < len(nums); i++ {
		splitSum += nums[i]

		mm := backTracking(nums, i+1, splitNum-1, sum, math.MaxInt64)
		minNum = min(minNum, max(splitNum, mm))
		if len(nums)-i+1 < splitNum {
			break
		}
	}

	return minNum
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(splitArray([]int{1, 4, 4}, 3))
	fmt.Println(splitArray([]int{4, 4}, 2))
}

// 输入：nums = [7,2,5,10,8], m = 2
// 输出：18
// 解释：
// 一共有四种方法将 nums 分割为 2 个子数组。 其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
// 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。/
