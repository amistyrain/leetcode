package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return (A != nil && B != nil) && (dp(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func dp(A, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return dp(A.Left, B.Left) && dp(A.Right, B.Right)
}
