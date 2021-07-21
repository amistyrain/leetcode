package main

import "fmt"

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}

	return result
}

// @lc code=end

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
