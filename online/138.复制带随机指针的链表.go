package main

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	oldIndex := make(map[*Node]int, 0)
	h := head
	k := 0
	for h != nil {
		oldIndex[h] = k
		h = h.Next
		k++
	}

	vals := make(map[int]*Node, k)
	newhead := &Node{}
	iter := newhead
	j := 0
	for head != nil {
		n := &Node{Val: head.Val}
		iter.Next = n
		if head.Random != nil {
			n.Random = &Node{Val: oldIndex[head.Random]}
		}
		iter = iter.Next
		vals[j] = n
		head = head.Next
		j++
	}

	newhead = newhead.Next
	iter = newhead
	for iter != nil {
		if iter.Random != nil {
			iter.Random = vals[iter.Random.Val]
		}
		iter = iter.Next
	}

	return newhead
}

// @lc code=end
//
// 返回复制链表的头节点。
//
// 用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：
//
// val：一个表示 Node.val 的整数。
// random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。
// 你的代码 只 接受原链表的头节点 head 作为传入参数。
