package medium

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var test func(i, j, k int) bool
	test = func(i, j, k int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '0' {
			return false
		}
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return board[i][j] == word[k]
		}
		temp := board[i][j]
		board[i][j] = '0'
		if test(i-1, j, k+1) || test(i, j-1, k+1) || test(i+1, j, k+1) || test(i, j+1, k+1) {
			return true
		}
		board[i][j] = temp
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if test(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
