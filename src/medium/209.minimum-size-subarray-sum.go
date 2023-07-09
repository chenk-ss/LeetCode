package medium

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	length := len(nums) + 1
	i, sum := 0, 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			if j-i+1 < length {
				length = j - i + 1
			}
			sum -= nums[i]
			i++
		}
	}
	if length == len(nums)+1 {
		return 0
	}
	return length
}

// @lc code=end
