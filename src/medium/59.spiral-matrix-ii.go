package medium

/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}
	left, top, right, buttom := 0, 0, n-1, n-1
	count := 1
	for {
		for i := left; i <= right; i++ {
			result[top][i] = count
			count++
		}
		top++
		if top > buttom {
			break
		}

		for i := top; i <= buttom; i++ {
			result[i][right] = count
			count++
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			result[buttom][i] = count
			count++
		}
		buttom--
		if buttom < top {
			break
		}

		for i := buttom; i >= top; i-- {
			result[i][left] = count
			count++
		}
		left++
		if left > right {
			break
		}
	}
	return result
}

// @lc code=end
