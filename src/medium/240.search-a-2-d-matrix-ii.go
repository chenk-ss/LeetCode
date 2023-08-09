package medium

/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	height, width := len(matrix), len(matrix[0])
	for i := 0; i < height; i++ {
		if matrix[i][0] > target {
			return false
		}
		if matrix[i][width-1] < target {
			continue
		}
		if matrix[i][0] == target || matrix[i][width-1] == target {
			return true
		}
		for j := 0; j < width; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// @lc code=end
