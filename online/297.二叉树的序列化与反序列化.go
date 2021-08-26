package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

const SEQ = `,`
const NIL = `#`

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var traverse func(*TreeNode)
	str := []string{}

	traverse = func(root *TreeNode) {
		if root == nil {
			str = append(str, NIL)
			str = append(str, SEQ)
			return
		}

		str = append(str, strconv.Itoa(root.Val))
		str = append(str, SEQ)

		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)

	return strings.Join(str, ``)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	nodes := strings.Split(data, SEQ)
	if len(nodes) == 0 {
		return nil
	}

	var retarverse func(*[]*TreeNode) *TreeNode

	treeNodes := make([]*TreeNode, 0, len(nodes)-1)
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i] == NIL {
			treeNodes = append(treeNodes, nil)
		} else {
			val, _ := strconv.Atoi(nodes[i])
			treeNodes = append(treeNodes, &TreeNode{Val: val})
		}

	}

	retarverse = func(nodes *[]*TreeNode) *TreeNode {
		if len(*nodes) == 0 {
			return nil
		}

		n := (*nodes)[0]
		*nodes = (*nodes)[1:]
		if n == nil {
			return nil
		}

		n.Left = retarverse(nodes)
		n.Right = retarverse(nodes)

		return n

	}

	return retarverse(&treeNodes)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end
