package medium

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	result[0] = searchRangeFunc(nums, target, false)
	result[1] = searchRangeFunc(nums, target, true)
	return result
}

func searchRangeFunc(nums []int, target int, flag bool) int {
	left, right, mid := 0, len(nums)-1, 0
	result := -1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			result = mid
			if flag {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return result
}

// @lc code=end
