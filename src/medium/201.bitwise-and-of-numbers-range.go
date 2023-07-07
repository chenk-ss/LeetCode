package medium

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	count := 0
	for left != right {
		left >>= 1
		right >>= 1
		count++
	}
	return left << count
}

// @lc code=end
