/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	p := head
	q := head

	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	newNode := reverseList(q, p)
	q.Next = reverseKGroup(p, k)

	return newNode
}
func reverseList(head, n *ListNode) *ListNode {
	if head == nil {
		return head
	}

	q := head
	p := head.Next
	q.Next = nil
	r := head

	for p != n {
		r = p.Next
		p.Next = q
		q = p
		p = r
	}

	return q
}

// @lc code=end

