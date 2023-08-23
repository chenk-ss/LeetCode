package medium

/*
 * @lc app=leetcode id=396 lang=golang
 *
 * [396] Rotate Function
 */

// @lc code=start
func maxRotateFunction(nums []int) int {
	length, max, sum := len(nums), 0, 0
	if length == 0 {
		return 0
	}
	for i, v := range nums {
		max += i * v
		sum += v
	}
	cur := max
	for i := length - 1; i >= 1; i-- {
		// cur = cur + (sum - nums[i]) - (length-1)*nums[i]
		cur = cur + sum - length*nums[i]
		if cur > max {
			max = cur
		}
	}
	return max
}

// @lc code=end
