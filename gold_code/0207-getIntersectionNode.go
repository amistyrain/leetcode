package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	a, b := headA, headB
	for a != nil {
		lenA++
		a = a.Next
	}
	for b != nil {
		lenB++
		b = b.Next
	}

	if lenA > lenB {
		i := lenA - lenB
		for i > 0 {
			headA = headA.Next
			i--
		}
	} else {
		i := lenB - lenA
		for i > 0 {
			headB = headB.Next
			i--
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
