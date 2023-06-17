package easy

/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for _, v := range nums {
		if v != 1 {
			count = 0
			continue
		}
		count++
		if count > max {
			max = count
		}
	}
	return max
}

// @lc code=end
