package medium

/*
 * @lc app=leetcode id=918 lang=golang
 *
 * [918] Maximum Sum Circular Subarray
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	max, min, maxSum, minSum, totalSum := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max = max918(nums[i], max+nums[i])
		min = min918(nums[i], min+nums[i])
		maxSum = max918(maxSum, max)
		minSum = min918(minSum, min)
		totalSum += nums[i]
	}
	if totalSum == minSum {
		return maxSum
	}
	return max918(maxSum, totalSum-minSum)
}

func max918(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min918(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
