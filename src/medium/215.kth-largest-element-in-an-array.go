package medium

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	k = k - 1
	for left < right {
		index := findKthLargestFunc(nums, left, right)
		if index < k {
			left = index + 1
		} else if index > k {
			right = index - 1
		} else {
			break
		}
	}
	return nums[k]
}

func findKthLargestFunc(nums []int, i, j int) int {
	pivot := nums[j]
	left := i
	for a := i; a < j; a++ {
		if nums[a] > pivot {
			nums[left], nums[a] = nums[a], nums[left]
			left++
		}
	}
	nums[j], nums[left] = nums[left], nums[j]
	return left
}

// @lc code=end
