package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxNum := make([]int, len(nums))
	minNum := make([]int, len(nums))
	maxNum[0] = nums[0]
	minNum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		maxNum[i] = max(nums[i]*maxNum[i-1], max(nums[i]*minNum[i-1], nums[i]))
		minNum[i] = min(nums[i]*maxNum[i-1], min(nums[i]*minNum[i-1], nums[i]))
	}

	res := maxNum[0]
	for i := 1; i < len(maxNum); i++ {
		res = max(res, maxNum[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// 输入: [2,3,-2,4]
// 输出: 6
// 解释: 子数组 [2,3] 有最大乘积 6。

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4, 5, -6}))
}

// 2
// 2*3
// -12 -6
// -48
//20
// -240
