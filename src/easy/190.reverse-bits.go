package easy

/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	result := uint32(0)
	for i := 0; i < 32; i++ {
		result = result<<1 + num&1
		num = num >> 1
	}
	return result
}

// @lc code=end
