/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(tree *TreeNode) []int {
	if tree == nil {
		return nil
	}
	res := []int{}
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

		res = append(res, tmp[len(tmp)-1])
	}

	return res
}

// @lc code=end

