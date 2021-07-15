package main

import "fmt"

// 自顶向下
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

// 转换成01背包
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum || (target+sum)%2 != 0 {
		return 0
	}
	target = (target + sum) / 2

	coins := nums
	amount := target

	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= len(coins); i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i][j] + dp[i-1][j-coins[i-1]]
			}
		}
	}

	return dp[len(coins)][amount]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3))
}
