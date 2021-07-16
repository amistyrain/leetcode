package main

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack  []int
	helper []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0),
		helper: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	if len(this.helper) == 0 {
		this.helper = append(this.helper, val)
	} else if this.helper[len(this.helper)-1] > val {
		this.helper = append(this.helper, val)
	} else {
		this.helper = append(this.helper, this.helper[len(this.helper)-1])
	}

	return
}

func (this *MinStack) Pop() {
	if len(this.helper) == 0 {
		return
	}
	this.helper = this.helper[:len(this.helper)-1]
	this.stack = this.stack[:len(this.stack)-1]

	return
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}

	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.helper) == 0 {
		return 0
	}

	return this.helper[len(this.helper)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
