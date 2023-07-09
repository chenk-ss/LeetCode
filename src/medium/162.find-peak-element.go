package medium

/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	return findPeakElementFunc(nums, 0, len(nums)-1)
}

func findPeakElementFunc(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if (mid == 0 && nums[mid] > nums[mid+1]) ||
		(mid == len(nums)-1 && nums[mid] > nums[mid-1]) ||
		(mid != 0 && mid != len(nums)-1 && nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1]) {
		return mid
	}
	if left == right {
		return -1
	}
	res := findPeakElementFunc(nums, left, mid-1)
	if res != -1 {
		return res
	}
	res = findPeakElementFunc(nums, mid+1, right)
	return res
}

// @lc code=end
