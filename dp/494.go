package main

import "fmt"

// 自顶向下
// func findTargetSumWays(nums []int, target int) int {
// 	var dp func(i, target int)

// 	res := 0

// 	dp = func(i, target int) {
// 		if target == 0 && i == len(nums) {
// 			res += 1
// 			return
// 		}
// 		if i >= len(nums) {
// 			return
// 		}
// 		dp(i+1, target-nums[i])
// 		dp(i+1, target-nums[i]*-1)

// 		return
// 	}
// 	dp(0, target)

// 	return res
// }

func findTargetSumWays(nums []int, target int) int {
	var dp func(i, target int)

	res := 0

	dp = func(i, target int) {
		if target == 0 && i == len(nums) {
			res += 1
			return
		}
		if i >= len(nums) {
			return
		}
		dp(i+1, target-nums[i])
		dp(i+1, target-nums[i]*-1)

		return
	}
	dp(0, target)

	return res
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
