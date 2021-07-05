package main

import "fmt"

// 双指针
// func isSubsequence(s string, t string) bool {
// 	if len(s) == 0 {
// 		return true
// 	}
// 	i := len(s) - 1
// 	j := len(t) - 1

// 	for j >= 0 {
// 		if s[i] == t[j] {
// 			if i == 0 {
// 				return true
// 			}
// 			j--
// 			i--
// 			continue
// 		}
// 		j--
// 	}

// 	return false
// }

func main() {
	fmt.Println(isSubsequence(`bb`, `ab`))
}

func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
