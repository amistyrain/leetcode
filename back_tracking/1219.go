package main

import "fmt"

// 24
// [[0,6,0],
//  [5,8,7],
//  [0,9,0]]
func getMaximumGold(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	maximumGold := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			maximumGold = max(maximumGold, dfs(grid, i, j))
		}
	}

	return maximumGold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(grid [][]int, row, col int) int {
	if col < 0 || col >= len(grid[0]) || row < 0 || row >= len(grid) || grid[row][col] == 0 {
		return 0
	}
	tmp := grid[row][col]
	grid[row][col] = 0

	up := dfs(grid, row-1, col)
	down := dfs(grid, row+1, col)
	left := dfs(grid, row, col-1)
	right := dfs(grid, row, col+1)

	grid[row][col] = tmp
	maxnum := max(up, max(down, max(left, right)))

	return maxnum + grid[row][col]
}

func main() {
	fmt.Println(getMaximumGold([][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}))
}
