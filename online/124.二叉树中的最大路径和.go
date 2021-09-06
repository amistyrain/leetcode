package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	ans = math.MinInt64
	dp(root)
	return ans
}

var ans int

func dp(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(dp(root.Left), 0)
	right := max(dp(root.Right), 0)

	ans = max(ans, left+right+root.Val)

	return root.Val + max(left, right)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
