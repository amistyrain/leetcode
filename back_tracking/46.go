package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	res = [][]int{}
	backtrack(nums, make([]int, 0))
	fmt.Println(res)
	return res
}

var res [][]int

func backtrack(nums []int, track []int) {
	if len(nums) == len(track) {
		// NOTE 引用坑
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for _, v := range nums {
		if contains(track, v) {
			continue
		}
		track = append(track, v)

		backtrack(nums, track)

		track = track[:len(track)-1]
	}

	return
}

func contains(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}

	return false
}

func main() {
	permute([]int{5, 6, 4, 2})
	fmt.Println(res)
}
