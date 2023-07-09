package medium

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	left, right, mid := 0, len(nums)-1, 0
	for left < right {
		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// @lc code=end
