/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			left := inorder[0:i]
			root.Left = buildTree(preorder[1:len(left)+1], left)
			root.Right = buildTree(preorder[len(left)+1:], inorder[i+1:])
		}
	}

	return root
}

// @lc code=end

