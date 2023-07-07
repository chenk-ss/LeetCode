package medium

/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	maxSide, m, n := 0, len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = maximalSquareFunc(maximalSquareFunc(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
				dp[i][j]++
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}
	return maxSide * maxSide
}

func maximalSquareFunc(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
