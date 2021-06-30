package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 4, 6, 2, 2, 7}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		dp[i+2] = max(dp[i-1+2], dp[i-2+2]+nums[i])
	}

	return dp[len(nums)+2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b

}
