package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(intervalsSort(intervals))
	res := [][]int{}
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		} else {
			res = append(res, interval)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type intervalsSort [][]int

func (inter intervalsSort) Len() int {
	return len(inter)
}

// 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
func (inter intervalsSort) Less(i, j int) bool {
	return inter[i][0] < inter[j][0]
}

// 交换 i 和 j 索引的两个元素的位置
func (inter intervalsSort) Swap(i, j int) {
	inter[i], inter[j] = inter[j], inter[i]
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
}
