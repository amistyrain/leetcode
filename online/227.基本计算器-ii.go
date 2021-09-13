package main

import "strconv"

/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start
func calculate(s string) int {
	ans := 0
	stack := []int{}
	num := 0
	sign := '+'

	for i, v := range s {
		isdigits := v >= '0' && v <= '9'
		if isdigits {
			val, _ := strconv.Atoi(string(v))
			num = num*10 + val
		}
		if !isdigits && v != ' ' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			sign = v
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}

	return ans
}

// @lc code=end
