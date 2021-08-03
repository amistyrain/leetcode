package main

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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

var idx int
var result int

func kthSmallest(root *TreeNode, k int) int {
	idx = 0
	result = 0
	dp(root, k)

	return result
}

func dp(root *TreeNode, k int) {
	if root == nil {
		return
	}
	kthSmallest(root.Left, k)
	idx++
	if idx == k {
		result = root.Val
		return
	}

	kthSmallest(root.Right, k)

}

// @lc code=end
