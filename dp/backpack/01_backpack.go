package main

import "fmt"

func main() {
	wt := []int{2, 1, 3}
	vals := []int{4, 2, 3}
	N := 3
	W := 4
	fmt.Println(maxValue(wt, vals, N, W))
}

func maxValue(wt, vals []int, N, W int) int {
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w-wt[i-1]]+vals[i-1], dp[i-1][w])
			}
		}
	}

	return dp[N][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// dp[i][w]
