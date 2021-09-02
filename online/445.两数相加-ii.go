package main

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)
	for l1 != nil {
		stack1 = append(stack1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2)
		l2 = l2.Next
	}
	carry := 0

	head := &ListNode{}
	head = nil
	for len(stack1) != 0 || len(stack2) != 0 || carry != 0 {
		v1, v2 := 0, 0
		if len(stack1) != 0 {
			v1 = stack1[len(stack1)-1].Val
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			v2 = stack2[len(stack2)-1].Val
			stack2 = stack2[:len(stack2)-1]
		}
		v := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		node := &ListNode{Val: v, Next: head}
		head = node
	}

	return head
}

// @lc code=end
