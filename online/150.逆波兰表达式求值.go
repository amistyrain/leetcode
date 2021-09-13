package main

import "strconv"

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			n := stack[len(stack)-2] + stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], n)
		case "-":
			n := stack[len(stack)-2] - stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], n)
		case "*":
			n := stack[len(stack)-2] * stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], n)
		case "/":
			n := stack[len(stack)-2] / stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], n)
		default:
			n, _ := strconv.Atoi(tokens[i])
			stack = append(stack, n)
		}
	}

	return stack[0]
}

// @lc code=end
