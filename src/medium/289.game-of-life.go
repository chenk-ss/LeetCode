package medium

/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	var gameOfLifeFuncCheck func(i, j int) int
	gameOfLifeFuncCheck = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 0 {
			return 0
		}
		return 1
	}

	var gameOfLifeFunc func(i, j int) int
	gameOfLifeFunc = func(i, j int) int {
		res := 0
		res += gameOfLifeFuncCheck(i-1, j-1)
		res += gameOfLifeFuncCheck(i-1, j)
		res += gameOfLifeFuncCheck(i-1, j+1)
		res += gameOfLifeFuncCheck(i, j-1)
		res += gameOfLifeFuncCheck(i, j+1)
		res += gameOfLifeFuncCheck(i+1, j-1)
		res += gameOfLifeFuncCheck(i+1, j)
		res += gameOfLifeFuncCheck(i+1, j+1)
		if board[i][j] == 1 && res >= 2 && res <= 3 {
			dp[i][j] = 1
		}
		if board[i][j] == 0 && res == 3 {
			dp[i][j] = 1
		}
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			gameOfLifeFunc(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = dp[i][j]
		}
	}
}

// @lc code=end
