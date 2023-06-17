package easy

/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	for n > 1 {
		if n&1 != 0 {
			return false
		}
		n = n >> 1
	}
	return n == 1
}

// @lc code=end
