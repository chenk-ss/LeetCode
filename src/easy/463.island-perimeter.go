package easy

/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	num1, numMix := 0, 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				num1++
				if j != len(grid[i])-1 && grid[i][j+1] == 1 {
					numMix++
				}
				if i != len(grid)-1 && grid[i+1][j] == 1 {
					numMix++
				}
			}
		}
	}
	return 4*num1 - 2*numMix
}

// @lc code=end
