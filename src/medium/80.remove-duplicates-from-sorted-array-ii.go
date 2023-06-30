package medium

/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func removeDuplicates2(nums []int) int {
	left := 0
	count := 1
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == last {
			if count < 2 {
				left++
				nums[left] = nums[i]
			}
			count++
		} else {
			left++
			nums[left] = nums[i]
			count = 1
			last = nums[i]
		}
	}
	return left + 1
}

// @lc code=end
