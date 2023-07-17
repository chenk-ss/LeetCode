package medium

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return rob213Max(rob213Func(nums[1:]), rob213Func(nums[:len(nums)-1]))
}

func rob213Func(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	list := make([]int, len(nums))
	list[0] = nums[0]
	list[1] = nums[1]
	if nums[1] < nums[0] {
		list[1] = nums[0]
	}
	for i := 2; i < len(nums); i++ {
		list[i] = rob213Max(list[i-1], list[i-2]+nums[i])
	}
	return list[len(nums)-1]
}

func rob213Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
