package main

import (
	"fmt"
	"strings"
)

func permutation(s string) []string {
	res = []string{}
	existMap := make(map[int]bool)
	backtrack(existMap, s, ``)
	return res
}

var res []string

func backtrack(existMap map[int]bool, s string, track string) {
	if len(s) == len(track) {
		res = append(res, track)
		return
	}

	for i, v := range s {
		if strings.Contains(track, string(v)) {
			continue
		}
		track += string(v)

		backtrack(existMap, s, track)

		track = track[:len(track)-1]
	}

	return
}

func main() {
	permutation(`abc`)
	fmt.Println(res)
}
