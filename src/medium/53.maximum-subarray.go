package medium

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	left, right, max, sum := 0, 0, nums[0], 0
	for left <= right && left < len(nums) && right < len(nums) {
		sum += nums[right]
		if sum > max {
			max = sum
		}
		if sum < 0 && right+1 < len(nums) {
			left = right + 1
			sum = 0
		}
		right++
	}
	return max
}

// @lc code=end
