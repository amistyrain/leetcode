package main

import "fmt"

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

// 1->2
// 1->2->3->2->1
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}
	// 翻转
	mid := reverList(slow.Next)

	for mid != nil {
		if mid.Val != head.Val {
			return false
		}
		mid = mid.Next
		head = head.Next
	}

	return true
}

func reverList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p := head
	head = head.Next
	p.Next = nil

	for head != nil {
		tmp := head.Next
		head.Next = p
		p = head
		head = tmp
	}

	return p
}

func main() {
	fmt.Println(isPalindrome(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	))
	fmt.Println(isPalindrome(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	))
	fmt.Println(isPalindrome(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	))
}
