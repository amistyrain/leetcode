package main

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	queue  []int
	helper []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue:  make([]int, 0),
		helper: make([]int, 0),
	}
}

// 1 2
/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.helper = append(this.helper, x)
	this.helper = append(this.helper, this.queue...)
	tmp := this.helper
	this.helper = []int{}
	this.queue = tmp
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if len(this.queue) == 0 {
		return 0
	}
	x := this.queue[0]
	this.queue = this.queue[1:]

	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.queue) == 0 {
		return 0
	}
	x := this.queue[0]

	return x
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
