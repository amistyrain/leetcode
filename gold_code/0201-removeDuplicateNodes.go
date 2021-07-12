package main

func main() {

}

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

// not sort list
func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	valsSet := make(map[int]struct{})
	valsSet[p.Val] = struct{}{}
	for p.Next != nil {
		val := p.Next.Val
		if _, ok := valsSet[val]; ok {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
		valsSet[val] = struct{}{}
	}

	return head
}

// 0 0 1 1 1 2 2 -> 0 1 2
func removeDuplicateSorteda(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil

	return head
}

// 1 2 2 3 3 3 4  5
func removeDuplicateSorted(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := &ListNode{}
	q := p
	for head != nil {
		if head.Next == nil || head.Val != head.Next.Val {
			p.Next = head
			p = p.Next
		}
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		head = head.Next

	}
	p.Next = nil

	return q.Next
}
