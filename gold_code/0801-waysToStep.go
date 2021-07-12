package main

import (
	"fmt"
)

func waysToStep(n int) int {
	if n < 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}

	return dp[n]
}

// 1 2 3 4 5 6 7 8 9
// 1 2 4 7 13

func main() {
	fmt.Println(waysToStep(5))
	fmt.Println(waysToStep(4))
	fmt.Println(waysToStep(3))
	fmt.Println(waysToStep(3))
	fmt.Println(waysToStep(76))
	// fmt.Println(waysToStep(63))
}
