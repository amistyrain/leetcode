package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	backtrack(candidates, target, []int{}, 0, 0)
	return res
}

var res [][]int

func backtrack(candidates []int, target int, track []int, sum int, start int) {
	if sum > target {
		return
	}
	if sum == target {
		res = append(res, append([]int{}, track...))
		return
	}
	for i := start; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backtrack(candidates, target, track, sum, i)

		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}

func main() {
	combinationSum([]int{2, 3, 8}, 8)
	fmt.Println(res)
}
