package main

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
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

func rotateRight(head *ListNode, k int) *ListNode {
	if k < 1 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	l1 := head
	for l1.Next != nil {
		n++
		l1 = l1.Next
	}
	k = n - k%n
	l1.Next = head

	for k > 0 {
		l1 = l1.Next
		k--
	}
	newHead := l1.Next
	l1.Next = nil

	return newHead
}

// @lc code=end
