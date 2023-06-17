package easy

/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
func findComplement(num int) int {
	count, copy := 0, num
	for copy > 0 {
		count = count<<1 + 1
		copy = copy >> 1
	}
	return count ^ num
}

// @lc code=end
