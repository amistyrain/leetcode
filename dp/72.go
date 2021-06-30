package main

import "fmt"

// 自顶向下
// func minDistance(word1 string, word2 string) int {
// 	dpTable := make(map[string]int)
// 	i := len(word1) - 1
// 	j := len(word2) - 1
// 	var dp func(int, int) int
// 	dp = func(i, j int) int {
// 		if v, ok := dpTable[dpTableKey(i, j)]; ok {
// 			return v
// 		}
// 		if i == -1 {
// 			return j + 1
// 		}
// 		if j == -1 {
// 			return i + 1
// 		}

// 		if word1[i] == word2[j] {
// 			dpTable[dpTableKey(i, j)] = dp(i-1, j-1)
// 		} else {
// 			dpTable[dpTableKey(i, j)] = min(
// 				dp(i-1, j)+1,
// 				min(dp(i, j-1)+1, dp(i-1, j-1)+1))
// 		}

// 		return dpTable[dpTableKey(i, j)]
// 	}

// 	return dp(i, j)
// }

// func dpTableKey(i, j int) string {
// 	return string(i) + `:` + string(j)
// }

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance(`intention`, `execution`))
}
