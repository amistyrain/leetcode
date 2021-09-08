package main

import "fmt"

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	for i := 0; i < n; i++ {
		j := i
		count := 0
		for ; j <= i+n; j++ {
			count = count + gas[j%n] - cost[j%n]
			if count < 0 {
				break
			}
		}
		if j == n+i+1 {
			return i
		}
	}

	return -1
}

// @lc code=end

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}

// 1 2 3 4 5 1 2 3 4 5
// 3 4 5 1 2 3 4 5 1 2
// 输入:
// gas  = [1,2,3,4,5]
// cost = [3,4,5,1,2]
//
// 输出: 3
//
// 解释:
// 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
// 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
// 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
// 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
// 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
// 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
// 因此，3 可为起始索引。
