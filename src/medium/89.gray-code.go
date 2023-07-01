package medium

/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 */

// @lc code=start
func grayCode(n int) []int {
	result := []int{}
	for i := 0; i < 1<<n; i++ {
		result = append(result, i^(i>>1))
	}
	return result
}

// @lc code=end
