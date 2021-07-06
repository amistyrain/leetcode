package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := treeheight(root.Left)
	right := treeheight(root.Right)
	if left-right > 1 || right-left > 1 {
		return false
	}
	return true
}

func treeheight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := treeheight(root.Left)
	right := treeheight(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}
