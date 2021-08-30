package main

import "fmt"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	strs := make([]string, min(numRows, len(s)))
	rowsNum := 0
	down := false
	for i := 0; i < len(s); i++ {
		strs[rowsNum] += string(s[i])
		if rowsNum == 0 || rowsNum == numRows-1 {
			down = !down
		}
		if down {
			rowsNum += 1
		} else {
			rowsNum -= 1
		}
	}
	ans := ``
	for i := 0; i < len(strs); i++ {
		ans += strs[i]
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// @lc code=end

func main() {
	fmt.Println(convert(`PAYPALISHIRING`, 4))
}
