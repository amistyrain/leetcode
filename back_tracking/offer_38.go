package main

import (
	"fmt"
	"strings"
)

func permutation(s string) []string {
	backtrack(s, ``)
	return res
}

var res []string

func backtrack(s string, track string) {
	if len(s) == len(track) {
		res = append(res, track)
		return
	}

	for _, v := range s {
		if strings.Contains(track, string(v)) {
			continue
		}
		track += string(v)

		backtrack(s, track)

		track = track[:len(track)-1]
	}

	return
}

func main() {
	permutation(`abc`)
	fmt.Println(res)
}
