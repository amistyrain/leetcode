package main

// 337 打家劫舍
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(tree *TreeNode) int {
	res := dfs(tree)
	return max(res[0], res[1])
}

func dfs(tree *TreeNode) []int {
	if tree == nil {
		return []int{0, 0}
	}
	l := dfs(tree.Left)
	r := dfs(tree.Right)
	selectT := tree.Val + l[1] + r[1]
	notSelectT := max(l[0], l[1]) + max(r[0], r[1])

	return []int{selectT, notSelectT}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
