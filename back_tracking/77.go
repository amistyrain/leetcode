package main

import "fmt"

func combine(n int, k int) [][]int {
	res = [][]int{}

	backtrack(n, k, 1, []int{})
	return res
}

var res [][]int

func backtrack(n, k, start int, track []int) {
	if k == len(track) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}

func main() {
	combine(4, 2)
	fmt.Println(res)
}
