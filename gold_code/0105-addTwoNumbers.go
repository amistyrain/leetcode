package main

import (
	"fmt"
	"math"
)

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

// 示例：

// 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
// 输出：2 -> 1 -> 9，即912
// 进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?

// 示例：

// 输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
// 输出：9 -> 1 -> 2，即912

// 2 4 3
// 5 6 4

// 342
// 465

// 7 0 8

// 7 7 1

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	num1 := l1.Val
	i := 1
	for l1.Next != nil {
		num1 += int(math.Pow10(i)) * l1.Next.Val
		l1.Next = l1.Next.Next
		i++
	}
	num2 := l2.Val
	i = 1
	for l2.Next != nil {
		num2 += int(math.Pow10(i)) * l2.Next.Val
		l2.Next = l2.Next.Next
		i++
	}
	sum := int64(num1 + num2)

	p := &ListNode{}
	q := p
	for {
		i := sum % 10
		sum /= 10
		p.Next = &ListNode{Val: int(i)}
		p = p.Next
		if sum == 0 {
			break
		}
	}

	return q.Next
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	return nil
}

func main() {
	head := addTwoNumbers(
		&ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
