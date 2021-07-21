package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第 N 位数字
 */

// @lc code=start
func findNthDigit(n int) int {
	if n < 10 {
		return n
	} else if 10 <= n && n < 100 {
		if n%2 == 0 {
			if n/10%2 == 0 {

			} else {

			}
			return n / 10
		} else {
			if n/10%2 == 0 {
				return (n%10 + 10) / 2
			} else {
				return n % 10 / 2
			}
		}
	} else {
		m := n
		i := 0
		for n != 0 {
			n = n / 10
			i++
		}
		if m%(i+1) == 0 {
			return m / int(math.Pow10(i+1))
		}
		return findNthDigit(m / 10)
	}

	return 0
}

// @lc code=end

func main() {
	fmt.Println(findNthDigit(10))
	fmt.Println(findNthDigit(11))
	fmt.Println(findNthDigit(12))
	fmt.Println(findNthDigit(21))
	fmt.Println(findNthDigit(20))
	fmt.Println(findNthDigit(1000))
	fmt.Println(findNthDigit(1097))
}

// 示例 1：

// 输入：3
// 输出：3
// 示例 2：

// 输入：11
// 输出：0
// 解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。
// 31 33 35 37 39 41
// 11 13 15 17 19 21 23 25 27 29
//  0  1  2  3  4  5  6  7  8  9
