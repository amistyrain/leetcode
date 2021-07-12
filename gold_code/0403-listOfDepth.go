package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	res := []*ListNode{}
	list := []*TreeNode{tree}

	for len(list) > 0 {
		length := len(list)
		head := &ListNode{}
		cur := head

		for i := 0; i < length; i++ {
			tail := list[i]
			cur.Next = &ListNode{Val: tail.Val}
			cur = cur.Next

			if tail.Left != nil {
				list = append(list, tail.Left)
			}
			if tail.Right != nil {
				list = append(list, tail.Right)
			}
		}
		list = list[length:]

		res = append(res, head.Next)
	}

	return res
}
