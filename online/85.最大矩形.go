package main

/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	heights := make([][]int, len(matrix))

	result := 0
	for i := 0; i < len(matrix); i++ {
		heights[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 {
				heights[i][j] = 1
			} else {
				heights[i][j] = heights[i-1][j] + 1
			}
		}
		result = max(result, largestRectangleArea(heights[i]))
	}

	return result
}

// func maximalRectangle(matrix [][]byte) int {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return 0
// 	}
// 	n, m := len(matrix), len(matrix[0])
// 	heights := make([]int, m)
// 	heights1 := make([]int, m)
//
// 	result := 0
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < m; j++ {
// 			if matrix[i][j] == '0' {
// 				heights1[j] = 0
// 				continue
// 			}
// 			heights1[j] = heights[j] + 1
// 		}
// 		copy(heights, heights1)
//
// 		result = max(result, largestRectangleArea(heights1))
// 	}
//
// 	return result
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	// [2,1,5,6,2,3]
	// order stack ： -1 2
	// 出栈元素:
	// 其左边界 left 和 右边界 right 中间夹住的位置，就是它的最大面积

	stack := make([]int, 0, len(heights))

	heights = append([]int{-1}, heights...) // 头哨兵
	heights = append(heights, -1)           // 尾哨兵
	maxArea := 0

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[peek(stack)] > heights[i] {
			s, p := pop(stack)
			stack = s
			if area := (i - peek(stack) - 1) * heights[p]; area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

func peek(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack []int) ([]int, int) {
	p := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return stack, p
}

// @lc code=end
