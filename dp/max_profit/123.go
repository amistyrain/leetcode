package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		k := make([][]int, 3)
		for j := 0; j < 3; j++ {
			k[j] = make([]int, 2)
		}
		dp[i] = k
	}
	dp[0][2][0] = 0
	dp[0][2][1] = -prices[0]
	dp[0][1][0] = 0
	dp[0][1][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
	}

	return dp[len(prices)-1][2][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// T[i][2][0] = max(T[i-1][2][0],T[i-1][2][1]+price)
// T[i][2][1] = max(T[i-1][2][1],T[i-1][1][0]-price)
// T[i][1][0] = max(T[i-1][1][0],T[i-1][1][1]+price)
// T[i][1][1] = max(T[i-1][1][1],-price)

// T[i][0][0] = 0

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
