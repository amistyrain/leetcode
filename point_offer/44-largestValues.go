package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}

	list := []*TreeNode{tree}
	res := []int{tree.Val}
	for len(list) > 0 {
		length := len(list)
		ans := math.MinInt64
		for i := 0; i < length; i++ {
			tail := list[i]

			if tail.Left != nil {
				list = append(list, tail.Left)
				ans = max(ans, tail.Left.Val)
			}
			if tail.Right != nil {
				list = append(list, tail.Right)
				ans = max(ans, tail.Right.Val)
			}
		}
		list = list[length:]
		if ans != math.MinInt64 {
			res = append(res, ans)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
