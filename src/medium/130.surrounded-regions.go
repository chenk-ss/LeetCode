package medium

/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	var solve130Func func(i, j int)
	solve130Func = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 'X' || dp[i][j] {
			return
		}
		dp[i][j] = true
		solve130Func(i-1, j)
		solve130Func(i, j-1)
		solve130Func(i+1, j)
		solve130Func(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				if board[i][j] == 'O' {
					solve130Func(i, j)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !dp[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

// @lc code=end
