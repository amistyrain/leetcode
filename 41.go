package main

import "fmt"

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	newNums := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			newNums = append(newNums, nums[i])
		}
	}

	if len(newNums) == 0 {
		return 1
	}
	for i := 0; i < len(newNums); i++ {
		num := abs(newNums[i])
		if num <= len(newNums) && newNums[num-1] > 0 {
			newNums[num-1] = -newNums[num-1]
		}
	}
	for i := 0; i < len(newNums); i++ {
		if newNums[i] > 0 {
			return i + 1
		}
	}

	return len(newNums) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

// @lc code=end

func main() {
	// 1 2 3 4 5
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))     // 1
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))         // 2
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))             // 3
	fmt.Println(firstMissingPositive([]int{-1, -2, -60, 40, 43})) // 1
}
