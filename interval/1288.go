package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	res := 0
	sort.Sort(intervalsSort(intervals))
	fmt.Println(intervals)
	left := intervals[0][0]
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		inter := intervals[i]
		if left <= inter[0] && right >= inter[1] {
			res++
		}
		if inter[1] >= right && inter[0] <= right {
			right = inter[1]
		}
		if inter[0] > right {
			left = inter[0]
			left = inter[1]
		}

	}

	return len(intervals) - res
}

type intervalsSort [][]int

func (inter intervalsSort) Len() int {
	return len(inter)
}

// 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
func (inter intervalsSort) Less(i, j int) bool {
	if inter[i][0] == inter[j][0] {
		return inter[i][1] > inter[j][1]
	}
	return inter[i][0] < inter[j][0]
}

// 交换 i 和 j 索引的两个元素的位置
func (inter intervalsSort) Swap(i, j int) {
	inter[i], inter[j] = inter[j], inter[i]
}

func main() {
	fmt.Println(removeCoveredIntervals([][]int{{1, 4}, {3, 6}, {2, 8}, {2, 7}}))
}
