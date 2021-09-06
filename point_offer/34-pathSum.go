package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0)
	dfs(root, target)
	return ans
}

var ans [][]int
var path []int

func dfs(root *TreeNode, target int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	target -= root.Val
	if root.Left == nil && root.Right == nil && target == 0 {
		ans = append(ans, append([]int{}, path...))
	}

	dfs(root.Left, target)
	dfs(root.Right, target)

	path = path[:len(path)-1]
}
