package medium

/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	length := len(nums)
	lastPoint := length - 1
	for i := length - 2; i >= 0; i-- {
		if i+nums[i] >= lastPoint {
			lastPoint = i
		}
	}
	return lastPoint == 0
}

// @lc code=end
