package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 4 5 1 9
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	p := &ListNode{}
	p.Next = head
	q := p
	for head != nil {
		if head.Val == val {
			q.Next = head.Next
			break
		}
		q = q.Next
		head = head.Next
	}

	return p.Next
}
