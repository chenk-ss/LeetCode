package medium

/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	bitCounts := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			bitCounts[i] += (num >> i) & 1
		}
	}

	result := int32(0)
	for i := 0; i < 32; i++ {
		if bitCounts[i]%3 != 0 {
			result |= (1 << i)
		}
	}

	return int(result)
}

// @lc code=end
