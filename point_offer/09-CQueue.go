package main

type CQueue struct {
	stack  []int
	helper []int
}

func Constructor() CQueue {
	return CQueue{
		stack:  make([]int, 0),
		helper: make([]int, 0),
	}

}

func (this *CQueue) AppendTail(value int) {
	this.stack = append(this.stack, value)

}

func (this *CQueue) DeleteHead() int {
	if len(this.helper) != 0 {
		x := this.helper[len(this.helper)-1]
		this.helper = this.helper[:len(this.helper)-1]
		return x
	}
	if len(this.stack) == 0 {
		return -1
	}
	for i := len(this.stack) - 1; i >= 0; i-- {
		this.helper = append(this.helper, this.stack[i])
	}
	this.stack = []int{}
	x := this.helper[len(this.helper)-1]
	this.helper = this.helper[:len(this.helper)-1]

	return x
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
