package main

import "fmt"

func subsets(nums []int) [][]int {
	backtrack(nums, []int{}, 0)
	return res
}

func backtrack(nums, track []int, start int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])

		backtrack(nums, track, i+1)

		track = track[:len(track)-1]
	}
}

var res [][]int

func main() {
	subsets([]int{1, 2, 3})
	fmt.Println(res)
}
