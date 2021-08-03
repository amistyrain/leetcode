/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	for i := 0; i < len(inorder); i++ {
		if postorder[len(postorder)-1] == inorder[i] {
			left := inorder[:i]
			root.Left = buildTree(left, postorder[:len(left)])
			root.Right = buildTree(inorder[i+1:], postorder[len(left):len(postorder)-1])
		}
	}

	return root
}

// @lc code=end

