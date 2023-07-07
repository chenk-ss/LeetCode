package medium

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if triangle[i+1][j] > triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j+1]
			} else {
				triangle[i][j] += triangle[i+1][j]
			}
		}
	}
	return triangle[0][0]
}

// @lc code=end
