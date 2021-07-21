package main

import (
	"fmt"
	"strconv"
)

// dp[n] 表示第 n 个位置方案数量，第 [n] 位与区间 [0:n-1] 位必然可以组成 dp[n-1] 个方案，如果前两位 [n-1:n] 组成的数字在区间 10-26 之间，第 [n-1:n] 位与区间 [0:n-2] 位必然可以组成 dp[n-2]个方案，所以可以根据该思路写出动态规划方程

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	dp := make([]int, len(numStr)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(numStr); i++ {
		s2 := numStr[i-2 : i]
		if s2 >= `10` && s2 <= `25` {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(numStr)]
}

func main() {
	// 输入: 12258
	// 输出: 5
	// 解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
	fmt.Println(translateNum(12258))
}
