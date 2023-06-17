package easy

/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	count := 0
	for x > 0 || y > 0 {
		if x&1+y&1 == 1 {
			count++
		}
		x = x >> 1
		y = y >> 1
	}
	return count
}

// @lc code=end
