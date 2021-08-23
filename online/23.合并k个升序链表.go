package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
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

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	h := head
	n := 0
	lastFlag := make([]bool, len(lists))
	for n < len(lists) {
		minNode := &ListNode{Val: math.MaxInt64}
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				if !lastFlag[i] {
					n++
					lastFlag[i] = true
				}
				continue
			}
			minNode = minNodes(minNode, lists[i])
		}
		if minNode.Val == math.MaxInt64 {
			continue
		}

		head.Next = minNode
		head = head.Next

		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if minNode == lists[i] {
				lists[i] = lists[i].Next
				break
			}
		}
	}

	return h.Next
}

func minNodes(a, b *ListNode) *ListNode {
	if a.Val > b.Val {
		return b
	}

	return a
}

// @lc code=end
