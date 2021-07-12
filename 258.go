package main

import "fmt"

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	sum := 0

	for num != 0 {
		sum += num % 10
		num /= 10
		if num == 0 && sum >= 10 {
			num = sum
			sum = 0
		}
		if sum < 10 && num == 0 {
			break
		}
	}

	return sum
}

// @lc code=end

func main() {
	fmt.Println(addDigits(38))
}
