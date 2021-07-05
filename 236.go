package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	mq := newSlidingWindow()

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			mq.push(nums[i])
		} else {
			mq.push(nums[i])
			res = append(res, mq.max())
			mq.pop(nums[i-k+1])
		}
	}

	return res
}

type maxQueue interface {
	max() int
	pop(int)
	push(int)
}

var _ maxQueue = (*slidingWindow)(nil)

type slidingWindow struct {
	list []int
}

func newSlidingWindow() *slidingWindow {
	return &slidingWindow{
		list: []int{},
	}
}

func (s *slidingWindow) max() int {
	if len(s.list) == 0 {
		return 0
	}

	return s.list[0]
}

func (s *slidingWindow) push(num int) {
	for i := len(s.list) - 1; i >= 0 && s.list[i] < num; i-- {
		s.list = s.list[:len(s.list)-1]
	}
	s.list = append(s.list, num)
}

func (s *slidingWindow) pop(n int) {
	if len(s.list) != 0 && s.list[0] == n {
		s.list = s.list[1:]
	}

	return
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
