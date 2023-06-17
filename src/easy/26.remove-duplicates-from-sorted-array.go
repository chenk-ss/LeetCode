package easy

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	k := 1
	last := nums[0]
	for i := 1; i < length; i++ {
		element := nums[i]
		if last != element {
			nums[k] = element
			last = element
			k++
		}
	}
	return k
}

// @lc code=end
