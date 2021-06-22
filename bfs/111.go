package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMinDepth := minDepth(root.Left)
	rightMinDepth := minDepth(root.Right)

	if leftMinDepth == 0 || rightMinDepth == 0 {
		return leftMinDepth + rightMinDepth + 1
	}
	if rightMinDepth > leftMinDepth {
		return leftMinDepth + 1
	}

	return rightMinDepth + 1
}

// bfs
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	list := []*TreeNode{}
	list = append(list, root)

	for len(list) > 0 {
		size := len(list)
		for i := 0; i < size; i++ {
			pop := list[i]
			if pop.Left == nil && pop.Right == nil {
				return depth
			}
			if pop.Left != nil {
				list = append(list, pop.Left)
			}
			if pop.Right != nil {
				list = append(list, pop.Right)
			}
		}
		list = list[size:]
		depth++
	}

	return depth
}
