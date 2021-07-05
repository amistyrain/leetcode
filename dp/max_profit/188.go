package main

import "fmt"

func maxProfitUnlimit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(prices)-1][0]
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k <= 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return maxProfitUnlimit(prices)
	}
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		nums := make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			nums[j] = make([]int, 2)
		}
		dp[i] = nums
	}

	for i := 0; i < k+1; i++ {
		dp[0][i][0] = 0
		dp[0][i][1] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[len(prices)-1][k][0]
}

// T[i][k][0] = max(T[i-1][k][0],T[i-1][k][1]+price)
// T[i][k][1] = max(T[i-1][k][1],T[i-1][k-1][0]-price)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
