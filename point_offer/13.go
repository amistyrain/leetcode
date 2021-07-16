package main

import "fmt"

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(visited, m, n, k, 0, 0)
}

func dfs(visited [][]bool, m, n, k int, i, j int) int {
	if i >= m || j >= n || visited[i][j] || bitSum(i, j) > k {
		return 0
	}
	visited[i][j] = true

	return 1 + dfs(visited, m, n, k, i+1, j) + dfs(visited, m, n, k, i, j+1)

}

func bitSum(x, y int) int {
	sum := 0
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	for y != 0 {
		sum += y % 10
		y /= 10
	}

	return sum
}

func main() {
	fmt.Println(movingCount(2, 3, 1))
}

// 00 01 02
// 10 11 12
