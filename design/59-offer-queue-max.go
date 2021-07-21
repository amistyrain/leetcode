package main

type MaxQueue struct {
	// 普通队列
	queue []int
	// 双端队列
	helper []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		queue:  make([]int, 0),
		helper: make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.helper) == 0 {
		return -1
	}

	return this.helper[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	i := len(this.helper) - 1
	for i >= 0 {
		if this.helper[i] < value {
			i--
			continue
		}
		break
	}
	this.helper = this.helper[:i+1]
	this.helper = append(this.helper, value)

}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	tail := this.queue[0]
	this.queue = this.queue[1:]
	if len(this.helper) != 0 && this.helper[0] == tail {
		this.helper = this.helper[1:]
	}

	return tail
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
