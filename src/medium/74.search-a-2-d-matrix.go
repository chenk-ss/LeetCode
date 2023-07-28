package medium

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix74(matrix [][]int, target int) bool {
	top, buttom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	rowMid, colMid := 0, 0
	for top <= buttom {
		rowMid = (top + buttom + 1) / 2
		if matrix[rowMid][0] == target {
			return true
		}
		if matrix[rowMid][0] > target {
			buttom = rowMid - 1
		} else {
			top = rowMid
		}
		if top == buttom {
			for left <= right {
				colMid = (left + right) / 2
				if matrix[top][colMid] == target {
					return true
				} else if matrix[top][colMid] > target {
					right = colMid - 1
				} else {
					left = colMid + 1
				}
			}
			return false
		}
	}
	return false
}

// @lc code=end
