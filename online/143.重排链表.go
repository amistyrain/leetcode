package main

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
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

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	stack := make([]*ListNode, 0)
	slow1 := slow
	for slow1 != nil {
		stack = append(stack, slow1)
		slow1 = slow1.Next
	}
	last2 := slow

	newHead := head
	for len(stack) != 0 && slow != nil {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		tmp := head.Next
		head.Next = last
		head = last
		head.Next = tmp
		head = head.Next

		slow = slow.Next
	}
	last2.Next = nil
	head = newHead

	return
}

// @lc code=end
