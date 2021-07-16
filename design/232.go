package main

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	stack  []int
	helper []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack:  make([]int, 0),
		helper: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.helper) != 0 {
		x := this.helper[len(this.helper)-1]
		this.helper = this.helper[:len(this.helper)-1]
		return x
	}
	if len(this.stack) == 0 {
		return 0
	}
	for i := len(this.stack) - 1; i >= 0; i-- {
		this.helper = append(this.helper, this.stack[i])
	}
	this.stack = []int{}
	x := this.helper[len(this.helper)-1]
	this.helper = this.helper[:len(this.helper)-1]

	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.helper) != 0 {
		return this.helper[len(this.helper)-1]
	}
	if len(this.stack) == 0 {
		return 0
	}

	for i := len(this.stack) - 1; i >= 0; i-- {
		this.helper = append(this.helper, this.stack[i])
	}
	this.stack = []int{}

	return this.helper[len(this.helper)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0 && len(this.helper) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
