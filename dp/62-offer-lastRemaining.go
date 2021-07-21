package main

import "fmt"

// 约瑟夫环
func lastRemaining(n int, m int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + m) % i
	}

	return dp[n]
}

func main() {
	// 0 1 2 3 4
	// 0 1 3 4
	// 1 3 4
	// 1 3
	// 3
	fmt.Println(lastRemaining(5, 3))
	fmt.Println(lastRemaining(5, 1))
	fmt.Println(lastRemaining(10, 17))
}
