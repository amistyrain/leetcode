package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	res = []string{}
	existMap := make([]bool, len(s))
	tmp := []byte(s)
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})

	backtrack(existMap, 0, tmp, ``)
	return res
}

var res []string

func backtrack(existMap []bool, start int, s []byte, track string) {
	if len(s) == len(track) {
		res = append(res, track)
		return
	}

	for i, v := range s {
		if existMap[i] || i > 0 && !existMap[i-1] && s[i] == s[i-1] {
			continue
		}
		existMap[i] = true
		track += string(v)

		backtrack(existMap, i+1, s, track)

		track = track[:len(track)-1]
		existMap[i] = false
	}

	return
}

func main() {
	permutation(`abc`)
	fmt.Println(res)
	permutation(`aab`)
	fmt.Println(res)
}
