package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res = []string{}
	backtrack(n, n, ``)
	return res
}

var res []string

func backtrack(left, right int, track string) {
	if right < left {
		return
	}
	if right < 0 || left < 0 {
		return
	}

	if left == 0 && right == 0 {
		res = append(res, track)
		return
	}

	track += `(`
	backtrack(left-1, right, track)
	track = track[:len(track)-1]

	track += `)`
	backtrack(left, right-1, track)
	track = track[:len(track)-1]
}

func main() {
	generateParenthesis(3)
	fmt.Println(res)
}
