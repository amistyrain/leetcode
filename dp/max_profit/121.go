package main

import "fmt"

func maxProfit(prices []int) int {
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
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// T[i][k][0] = max(T[i-1][k][0],T[i-1][k][1]+price)
// T[i][k][1] = max(T[i-1][k][1],T[i-1][k-1][0]-price)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
