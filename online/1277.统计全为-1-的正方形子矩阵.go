package main

/*
 * @lc app=leetcode.cn id=1277 lang=golang
 *
 * [1277] 统计全为 1 的正方形子矩阵
 */

// @lc code=start
func countSquares(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	result := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				}
				result += dp[i][j]
			}
		}
	}

	return result
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
