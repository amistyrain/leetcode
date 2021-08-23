package main

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, len(matrix))
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				}
				result = max(result, dp[i][j])
			}
		}
	}

	return result * result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// @lc code=end
