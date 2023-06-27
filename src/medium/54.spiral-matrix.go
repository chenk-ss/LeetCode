package medium

/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	result := []int{}
	top, left, buttom, right := 0, 0, len(matrix)-1, len(matrix[0])-1
	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > buttom {
			break
		}

		for i := top; i <= buttom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[buttom][i])
		}
		buttom--
		if top > buttom {
			break
		}

		for i := buttom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return result
}

// @lc code=end
