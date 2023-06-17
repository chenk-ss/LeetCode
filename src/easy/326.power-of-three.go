package easy

/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 {
		return false
	}
	for n > 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return n == 3
}

// @lc code=end
