package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if i+j == n-1 {
				continue
			}
			matrix[i][j], matrix[n-j-1][n-i-1] = matrix[n-j-1][n-i-1], matrix[i][j]
		}
	}
	return
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)
	fmt.Println(matrix)
}

// 给定 matrix =
// [
//   [1,2,3],
//   [4,5,6], 1 0  2 1
//   [7,8,9]
// ],

// 原地旋转输入矩阵，使其变为:
// [
//   [7,4,1],
//   [8,5,2],
//   [9,6,3]
// ]

// 1  2  3  4
// 5  6  7  8
// 9  10 11 12
// 13 14 15 16
