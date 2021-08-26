/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(tree *TreeNode) [][]int {
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

	return res
}

// @lc code=end

