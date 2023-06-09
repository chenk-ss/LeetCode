package src

/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {
	result := 0
	for num > 0 {
		if num&1 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}

// @lc code=end
