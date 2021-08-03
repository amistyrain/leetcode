package main

/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 把二叉搜索树转换为累加树
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
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	return dp(root)
}

func dp(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dp(root.Right)
	sum += root.Val
	root.Val = sum
	dp(root.Left)

	return root
}

// @lc code=end
