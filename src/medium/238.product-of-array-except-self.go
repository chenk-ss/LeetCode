package medium

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	mu := 1
	for i := 0; i < len(nums); i++ {
		res[i] = mu
		mu *= nums[i]
	}
	mu = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= mu
		mu *= nums[i]
	}
	return res
}

// @lc code=end
