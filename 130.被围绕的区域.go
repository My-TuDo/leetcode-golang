/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */

// @lc code=start
func solve(board [][]byte) {
	// DFS : 优化，解决重复调用栈的问题
	rolw := len(board)
	clos := len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r >= rolw || r < 0 || c >= clos || c < 0 || board[r][c] != 'O' {
			return
		}
		board[r][c] = '#'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
		return
	}

	// 先解决所有没被包围的区域
	for r := 0; r < rolw; r++ {
		dfs(r, 0)
		dfs(r, clos-1)
	}
	for c := 1; c < clos-1; c++ {
		dfs(0, c)
		dfs(rolw-1, c)
	}

	for r := 0; r < rolw; r++ {
		for c := 0; c < clos; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
			if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}
	return
}

// @lc code=end

