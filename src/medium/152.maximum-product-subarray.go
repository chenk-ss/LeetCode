package medium

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := nums[0]
	product := 1
	for i := 0; i < len(nums); i++ {
		product *= nums[i]
		if product > max {
			max = product
		}
		if product == 0 {
			product = 1
		}
	}
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		product *= nums[i]
		if product > max {
			max = product
		}
		if product == 0 {
			product = 1
		}
	}
	return max
}

// @lc code=end
