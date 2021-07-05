package main

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	res = [][]string{}

	track := make([]string, n)
	for i := 0; i < n; i++ {
		track[i] = strings.Repeat(`.`, n)
	}

	backtrack(track, 0)

	return res
}

var res [][]string

func backtrack(track []string, row int) {
	if len(track) == row {
		tmp := make([]string, len(track))
		copy(tmp, track)

		res = append(res, tmp)
		return
	}
	n := len(track)

	for i := 0; i < n; i++ {
		if !isValid(track, row, i) {
			continue
		}
		tmpStr := []byte(track[row])
		tmpStr[i] = 'Q'
		track[row] = string(tmpStr)

		backtrack(track, row+1)

		tmpStr = []byte(track[row])
		tmpStr[i] = '.'
		track[row] = string(tmpStr)
	}

}

func isValid(track []string, row, col int) bool {
	n := len(track)
	// 检查列
	for i := 0; i < n; i++ {
		if track[i][col] == 'Q' {
			return false
		}
	}
	// 检查左上角
	for i, j := row, col; i >= 0 && j >= 0; {
		if track[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// 检查左上角
	for i, j := row, col; i >= 0 && j < n; {
		if track[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}

func main() {
	solveNQueens(8)
}
