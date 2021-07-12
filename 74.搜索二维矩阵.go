/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	for i, j := n-1, 0; i >= 0 && j < m; {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			j++
			continue
		}
		if target < matrix[i][j] {
			i--
			continue
		}
	}

	return false
}

// @lc code=end

