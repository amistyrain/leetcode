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
		} else {
			p = p.Next
			valsSet[val] = struct{}{}
		}
	}

	return head
}
