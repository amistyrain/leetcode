/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(tree *TreeNode) [][]int {
	if tree == nil {
		return nil
	}
	res := [][]int{}
	list := []*TreeNode{tree}
	level := 0

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
		if level%2 == 1 {
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
			}
		}

		res = append(res, tmp)
		level++
	}

	// for i := 0; i < len(res)/2; i++ {
	// 	res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	// }

	return res
}

// @lc code=end

