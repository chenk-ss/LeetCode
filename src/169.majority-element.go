package src

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	num, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
		}
		if num == nums[i] {
			count++
		} else {
			count--
		}
	}
	return num
}

// @lc code=end
