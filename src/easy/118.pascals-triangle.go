package easy

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		rowList := make([]int, i+1)
		rowList[0] = 1
		rowList[i] = 1
		for j := 1; j < i; j++ {
			rowList[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = rowList
	}
	return result
}

// @lc code=end
