package main

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
// func isMatch(s string, p string) bool {
// 	if len(p) == 0 {
// 		return len(s) == 0
// 	}
// 	firstMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
// 	if len(p) >= 2 && p[1] == '*' {
// 		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
// 	} else {
// 		return firstMatch && isMatch(s[1:], p[1:])
// 	}
// }

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 2; i <= len(p); i++ {
		dp[0][i] = p[i-1] == '*' && dp[0][i-2]
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j-1] || (fmatch(s, p, i, j-1) && dp[i][j+1])
			} else {
				dp[i+1][j+1] = fmatch(s, p, i, j) && dp[i][j]
			}
		}
	}

	return dp[len(s)][len(p)]
}

func fmatch(s, p string, i, j int) bool {
	return s[i] == p[j] || p[j] == '.'
}

// @lc code=end
