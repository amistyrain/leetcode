package main

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h := head
	l1, l2 := head, head.Next
	mid := l2
	for l1 != nil && l1.Next != nil && l1.Next.Next != nil {
		l1.Next = l1.Next.Next
		l1 = l1.Next
		l2.Next = l2.Next.Next
		l2 = l2.Next
	}
	l1.Next = mid

	return h
}

// @lc code=end

// 输入: 1->2->3->4->5->NULL
// 输出: 1->3->5->2->4->NULL
