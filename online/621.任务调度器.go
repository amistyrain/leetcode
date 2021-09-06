package main

/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	count := [26]int{}
	for i := 0; i < len(tasks); i++ {
		count[tasks[i]-'A']++
	}
	x := 0
	m := 0
	for i := 0; i < len(count); i++ {
		if x < count[i] {
			x = count[i]
			m = 1
		} else if x == count[i] {
			m++
		}
	}

	return max(len(tasks), (n+1)*(x-1)+m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
