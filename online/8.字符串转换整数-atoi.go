package main

import "math"

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	s = s[i:len(s)]
	ans := 0
	sign := 1
	for i, v := range s {
		if v >= '0' && v <= '9' {
			ans = ans*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if sign == 1 && ans > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && -ans < math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * ans
}

// @lc code=end
