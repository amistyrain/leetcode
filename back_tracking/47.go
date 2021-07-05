package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	sort.Ints(nums)
	existMap := make(map[int]bool)
	backtrack(nums, []int{}, existMap)
	return res
}

var res [][]int

func backtrack(nums []int, track []int, existMap map[int]bool) {
	if len(nums) == len(track) {
		res = append(res, append([]int{}, track...))
		return
	}

	for i, v := range nums {
		if existMap[i] || (i > 0 && nums[i] == nums[i-1] && existMap[i-1]) {
			continue
		}
		track = append(track, v)
		existMap[i] = true
		backtrack(nums, track, existMap)
		existMap[i] = false
		track = track[:len(track)-1]
	}

	return
}

func main() {
	permuteUnique([]int{1, 1, 2, 5})
	fmt.Println(res)
}
