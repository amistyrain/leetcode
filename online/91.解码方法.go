package main

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1

	for i := 0; i < len(s); i++ {
		if s[i] > '0' && s[i] <= '9' {
			dp[i+1] += dp[i]
		}
		if i >= 1 && s[i-1] != '0' && ((s[i-1]-'0')*10+(s[i]-'0') <= 26) {
			dp[i+1] += dp[i-1]
		}
	}

	return dp[len(s)]
}

// @lc code=end
