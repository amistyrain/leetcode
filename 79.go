package main

import "fmt"

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, i, j) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, i, j int) bool {
	if i >= 0 && j >= 0 && i < len(board) && j < len(board[0]) && board[i][j] != '1' {
		if len(word) == 0 {
			return true
		}
		if word[0] != board[i][j] {
			return false
		}
		if len(word) == 1 && board[i][j] == word[0] {
			return true
		}
		tmp := board[i][j]
		board[i][j] = '1'
		if dfs(board, word[1:], i, j+1) || dfs(board, word[1:], i, j-1) ||
			dfs(board, word[1:], i-1, j) || dfs(board, word[1:], i+1, j) {
			return true
		}

		board[i][j] = tmp
	}

	return false
}

// @lc code=end

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, `ABCCED`))
}
