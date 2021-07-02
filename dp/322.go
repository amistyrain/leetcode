package main

import (
	"fmt"
	"math"
)

// 自顶向下
func coinChange(coins []int, amount int) int {
	mem := make(map[int]int, amount)
	var dp func(int) int
	dp = func(amount int) int {
		if v, ok := mem[amount]; ok {
			return v
		}
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		res := math.MaxInt64
		for i := 0; i < len(coins); i++ {
			tmp := dp(amount - coins[i])
			if tmp == -1 {
				continue
			}
			mem[amount-coins[i]] = tmp
			res = min(res, tmp+1)
		}
		if res == math.MaxInt64 {
			return -1
		}

		return res
	}

	return dp(amount)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 自底向上
// func coinChange(coins []int, acount int) int {
// 	dp := make([]int, acount+1)
// 	for i := 1; i < len(dp); i++ {
// 		dp[i] = i + 1
// 	}

// 	for i := 1; i <= acount; i++ {
// 		for j := 0; j < len(coins); j++ {
// 			if i-coins[j] < 0 {
// 				continue
// 			}
// 			if dp[i-coins[j]] == i-coins[j]+1 {
// 				continue
// 			}
// 			dp[i] = min(dp[i], dp[i-coins[j]]+1)
// 		}
// 	}

// 	if dp[acount] == acount+1 {
// 		return -1
// 	}

// 	return dp[acount]
// }

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
