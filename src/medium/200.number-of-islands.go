package medium

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	res := 0

	var numIslandsFunc func(i, j int)
	numIslandsFunc = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		numIslandsFunc(i-1, j)
		numIslandsFunc(i, j-1)
		numIslandsFunc(i, j+1)
		numIslandsFunc(i+1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				numIslandsFunc(i, j)
			}
		}
	}
	return res
}

// @lc code=end
