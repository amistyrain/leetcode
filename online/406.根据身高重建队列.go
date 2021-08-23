package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.SliceStable(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})

	// 按照k值插入到index=k的地方，index之后的往后移动
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i])
		people[p[1]] = p
	}

	return people
}

// @lc code=end

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
