package medium

/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	col := make([]int, len(matrix[0]))
	row := make([]int, len(matrix))
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

// @lc code=end
