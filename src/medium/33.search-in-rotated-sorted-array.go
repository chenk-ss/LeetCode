package medium

/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {

	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] > nums[mid] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// @lc code=end
