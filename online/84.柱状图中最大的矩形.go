package main

import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
// 栈+哨兵
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

func main() {
	fmt.Println(largestRectangleArea([]int{2, 2, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}
