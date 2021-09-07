package main

/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	row := n
	col := n
	left := matrix[0][0]
	right := matrix[n-1][n-1]

	for left < right {
		mid := left + (right-left)/2
		count := find(matrix, mid, row, col)
		if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

func find(matrix [][]int, mid int, row, col int) int {
	i := row - 1
	j := 0
	count := 0
	for j < col && i >= 0 {
		if matrix[i][j] <= mid {
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count
}

// @lc code=end

// 示例 1：
//
// 输入：matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
// 输出：13
// 解释：矩阵中的元素为 [1,5,9,10,11,12,13,13,15]，第 8 小元素是 13
// 示例 2：
//
// 输入：matrix = [[-5]], k = 1
// 输出：-5
