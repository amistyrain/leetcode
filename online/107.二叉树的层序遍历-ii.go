/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(tree *TreeNode) [][]int {
	if tree == nil {
		return nil
	}
	res := [][]int{}
	list := []*TreeNode{tree}

	for len(list) > 0 {
		length := len(list)
		tmp := make([]int, 0, length)

		for i := 0; i < length; i++ {
			tail := list[i]
			tmp = append(tmp, tail.Val)

			if tail.Left != nil {
				list = append(list, tail.Left)
			}
			if tail.Right != nil {
				list = append(list, tail.Right)
			}
		}
		list = list[length:]

		res = append(res, tmp)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}

// @lc code=end

