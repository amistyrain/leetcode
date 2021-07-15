/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	n := len(matrix)
	m := len(matrix[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] || matrix[i][j] != 0 {
				continue
			}
			visited[i][j] = true
			for x := 0; x < m; x++ {
				if matrix[i][x] != 0 {
					matrix[i][x] = 0
					visited[i][x] = true
				}
			}
			for y := 0; y < n; y++ {
				if matrix[y][j] != 0 {
					matrix[y][j] = 0
					visited[y][j] = true
				}
			}
		}
	}
}
// @lc code=end

