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
		tmp := 0
		if i >= 2 {
			tmp = dp[i-2][0] - prices[i]
		} else {
			tmp = -prices[i]
		}
		dp[i][1] = max(dp[i-1][1], tmp)
	}

	return dp[len(prices)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1, 2, 4}))
}
