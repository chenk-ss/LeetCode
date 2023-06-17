package easy

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	primer := make([]int, 0)
	for i := 0; i < rowIndex+1; i++ {
		rowList := make([]int, i+1)
		rowList[0] = 1
		rowList[i] = 1
		for j := 1; j < i; j++ {
			rowList[j] = primer[j-1] + primer[j]
		}
		primer = rowList
	}
	return primer
}

// @lc code=end
