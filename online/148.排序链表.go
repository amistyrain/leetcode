package main

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
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

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	r := sortList(slow.Next)
	slow.Next = nil
	l := sortList(head)

	newHead := &ListNode{}

	return mergeList(newHead, l, r)
}

func mergeList(head, l, r *ListNode) *ListNode {
	p := head
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l == nil {
		p.Next = r
	}
	if r == nil {
		p.Next = l
	}

	return head.Next
}

// @lc code=end
