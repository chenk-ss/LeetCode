package medium

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
				continue
			}
			grid[i][j] = grid[i][j] + minPathSumMin(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func minPathSumMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
