package main

import "fmt"

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			wordLen := len(word)
			if i-wordLen >= 0 && s[i-wordLen:i] == word {
				dp[i] = dp[i] || dp[i-wordLen]
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end

func main() {
	fmt.Println(wordBreak(`leetet`, []string{`leet`, `et`}))
}
