package medium

/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	cp := [][]int{}
	for _, v := range matrix {
		cp = append(cp, append([]int{}, v...))
	}
	n := len(matrix)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = cp[n-1-j][i]
		}
	}
}

// @lc code=end
