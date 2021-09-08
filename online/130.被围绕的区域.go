package main

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n := len(board)
	m := len(board[0])

	for i := 0; i < m; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)

	}

	for i := 0; i < n; i++ {
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'H' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && board[i][j] == 'O' {
		board[i][j] = 'H'
		dfs(board, i+1, j)
		dfs(board, i, j+1)
		dfs(board, i-1, j)
		dfs(board, i, j-1)
	}
}

// @lc code=end

//
// 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。

// ["X","X","X","X"]
// ["X","O","O","X"]
// ["X","X","O","X"]
// ["X","O","X","X"]

// ["X","X","X","X"]
// ["X","X","X","X"]
// ["X","X","X","X"]
// ["X","O","X","X"]
