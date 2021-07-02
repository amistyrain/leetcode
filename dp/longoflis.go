package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{1, 5, 76, 2, 8, 7, 9, 6}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	res := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			res = max(res, dp[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
