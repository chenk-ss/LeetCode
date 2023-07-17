package medium

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob198(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if nums[1] < nums[0] {
		nums[1] = nums[0]
	}
	for i := 2; i < length; i++ {
		nums[i] = robMaxFunc(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[length-1]
}

func robMaxFunc(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
