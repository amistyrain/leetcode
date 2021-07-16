package main

import "fmt"

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	i, j := 0, 0
	n, m := len(matrix)-1, len(matrix[0])-1
	res := make([]int, 0, n*m)
	for i <= n && j <= m {
		for k := j; k <= m; k++ {
			res = append(res, matrix[i][k])
		}
		for k := i + 1; k <= n; k++ {
			res = append(res, matrix[k][m])
		}
		if i < n && j < m {
			for k := m - 1; k > j; k-- {
				res = append(res, matrix[n][k])
			}

			for k := n; k > i; k-- {
				res = append(res, matrix[k][j])
			}

		}
		i++
		n--
		m--
		j++
	}

	return res
}

// @lc code=end

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix))
}
