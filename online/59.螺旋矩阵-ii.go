package main

import "fmt"

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	length := n
	i, j := 0, 0
	n, m := n-1, n-1
	num := 1
	res := make([][]int, length)
	for x := 0; x < length; x++ {
		res[x] = make([]int, length)
	}

	for i <= n && j <= m {
		for k := j; k <= m; k++ {
			res[i][k] = num
			num++
		}
		for k := i + 1; k <= n; k++ {
			res[k][m] = num
			num++
		}
		if i < n && j < m {
			for k := m - 1; k > j; k-- {
				res[n][k] = num
				num++
			}

			for k := n; k > i; k-- {
				res[k][j] = num
				num++
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
	fmt.Println(generateMatrix(3))

}
